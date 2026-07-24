package main

import (
	"fmt"
	"net/http"
	"net/mail"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// each field in the Struct has a Rule to check validity (specified by tag `validate:xxx`)
type fieldMeta struct {
	value reflect.Value // "phm1605@gmail.com" OR "18202123" OR "231090000036"
	rule  string        // "email" OR "zipcode" OR "creditcard"
}

// NEW: to validate field's value according to "rule" before changing it
// if all good then returns "nil error"
func validate(v reflect.Value, rule string) error {
	if rule == "" {
		return nil
	}
	switch rule {
	case "email":
		if v.Kind() != reflect.String {
			return fmt.Errorf("email validator requires a string")
		}
		// try parsing Email using library
		_, err := mail.ParseAddress(v.String())
		if err != nil {
			return fmt.Errorf("invalid email address")
		}

	case "zipcode":
		if v.Kind() != reflect.Int {
			return fmt.Errorf("zipcode validator requires an int")
		}
		zip := v.Int()
		if zip < 10000 || zip > 99999 {
			return fmt.Errorf("invalid US ZIP code")
		}

	case "creditcard":
		if v.Kind() != reflect.String {
			return fmt.Errorf("creditcard validator requires a string")
		}
		s := strings.ReplaceAll(v.String(), " ", "")
		s = strings.ReplaceAll(s, "-", "")
		if len(s) < 13 || len(s) > 19 {
			return fmt.Errorf("invalid credit card number")
		}
		// check if each char is a number
		for _, r := range s {
			if !unicode.IsDigit(r) {
				return fmt.Errorf("invalid credit card number")
			}
		}

	default:
		return fmt.Errorf("unknown validator %q", rule)
	}
	// All good
	return nil
}

// put "value" into "v" OR put "value" into 1 element of the Slice
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	// all good
	return nil
}

// ptr is for whatever "*struct" we are injecting information into
// request example: "http://localhost:12345/search?l=golang&l=books&max=2&x=true"
func Unpack(req *http.Request, ptr interface{}) error {
	// after this statement "req.Form" will be iterable (containing all parameters)
	if err := req.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]fieldMeta) // fieldInfo: its value AND rule-name that has been checked
	// reflect.ValueOf(ptr) returns "*struct-value"
	v := reflect.ValueOf(ptr).Elem() // reflect.Value of that data-struct
	for i := 0; i < v.NumField(); i++ {
		// a reflect.StructField; with "name","type" and optional "tag" inside
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag    // a reflect.StructTag
		name := tag.Get("http") // get value of tag with `http:` prefix i.e. "l","max","x"
		// if no tag then we take lowercase of field name e.g. "labels"/"maxresults"/"exact"
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = fieldMeta{
			value: v.Field(i),
			rule:  tag.Get("validate"),
		}
	}

	// name="email", values=["xxx"]; name="zip", values=[2]; ...
	for name, values := range req.Form {
		meta, ok := fields[name] // metadata of that field
		// ignore weird params from Request
		if !ok {
			continue
		}
		f := meta.value
		for _, value := range values {
			// if l="x" & l="y" many times => accumulate
			if f.Kind() == reflect.Slice {
				// "f" is reflect.Value of "[]string"
				// => f.Type().Elem() is "string"
				elem := reflect.New(f.Type().Elem()).Elem() // addressable reflect.Value of String
				// set 1 element of Slice (to be appended) here
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				// NEW: validate value of this field with RULE
				if err := validate(elem, meta.rule); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else { // if max=12 & max=8 many times => take the last i.e. max=8
				if err := populate(f, value); err != nil { // inject "value" from Request into "f"
					return fmt.Errorf("%s: %v", name, err)
				}
				// NEW: validate value of this field with RULE
				if err := validate(f, meta.rule); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil // all good
}

func search(resp http.ResponseWriter, req *http.Request) {
	// NEW: our data struct to extract parameters must contain this sensible information
	// http: contains the shorter and more concise parameter name
	// validate: RULE name to CHECK value before injecting values to "data"
	var data struct {
		Email string `http:"email" validate:"email"`
		ZIP   int    `http:"zip" validate:"zipcode"`
		Card  string `http:"card" validate:"creditcard"`
	}

	// unpack request's parameters to that struct (variable "data" has that struct type)
	if err := Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ... rest of handler
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

func main() {
	http.HandleFunc("/search", search)
	fmt.Println("Listening on :12345")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic(err)
	}
}
