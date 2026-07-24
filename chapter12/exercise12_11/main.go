package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// handle reflect.Value => convert to String
func formatValue(v reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.String:
		return v.String(), nil

	case reflect.Int:
		return strconv.FormatInt(v.Int(), 10), nil

	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil

	default:
		return "", fmt.Errorf("unsupported kind %s", v.Type())
	}
}

// v: interface of a "struct" or "*struct"
func Pack(v interface{}) (string, error) {
	rv := reflect.ValueOf(v) // struct OR *struct
	// if *struct
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	// now "rv" has surely "rt=struct"
	rt := rv.Type()

	values := url.Values{} // to collect Parameters for HTTP request using "values.Add(name, value)"
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i) // ["golang", "cpp"]
		info := rt.Field(i)  // Labels []string `http:"l"` OR MaxResults int `http:"max"`
		// tag
		name := info.Tag.Get("http") // "l"
		if name == "" {
			name = strings.ToLower(info.Name) // "labels" if "l" tag not there
		}

		if field.Kind() == reflect.Slice {
			for i := 0; i < field.Len(); i++ {
				s, err := formatValue(field.Index(i))
				if err != nil {
					return "", fmt.Errorf("%s: %v", name, err)
				}
				values.Add(name, s)
			}
		} else {
			s, err := formatValue(field)
			if err != nil {
				return "", fmt.Errorf("%s: %v", name, err)
			}
			values.Add(name, s)
		}
	}
	// format to "l=golang&l=books&max=20&x=true"
	return values.Encode(), nil
}

func main() {
	data := struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}{
		Labels:     []string{"golang", "books"},
		MaxResults: 20,
		Exact:      true,
	}

	query, err := Pack(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(query)
}
