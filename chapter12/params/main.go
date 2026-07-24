package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

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

	// map "l"=>reflect.Value(<Labels>); "max"=>reflect.Value(<MaxResults>); "x"=>reflect.Value(<Exact>)
	// (if no tag we use lowercase of field name as key)
	fields := make(map[string]reflect.Value)
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
		fields[name] = v.Field(i)
	}

	// name="l", values=["golang","books"]; name="max", values=[2]; ...
	for name, values := range req.Form {
		f := fields[name] // reflect.Value of value of that field
		// ignore weird params from Request
		if !f.IsValid() {
			continue
		}
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
				f.Set(reflect.Append(f, elem))
			} else { // if max=12 & max=8 many times => take the last i.e. max=8
				if err := populate(f, value); err != nil { // inject "value" from Request into "f"
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil // all good
}

func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // set default

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
