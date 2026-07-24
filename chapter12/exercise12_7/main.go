package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

// wrapper around a Scanner; to return the current token (word)
type lexer struct {
	scan  scanner.Scanner
	token rune // the current token TYPE e.g. scanner.Ident
}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan() // put "scanner.Ident" into "lex.token" (inside scanner there is the token value e.g. "x_var")
}

func (lex *lexer) text() string {
	return lex.scan.TokenText() // "x_var"
}

// pop the current token out of Scanner (if we want same Type as what Scanner is holding) and move to the next
func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

// NEW: for streaming API
type Decoder struct {
	lex *lexer // scanner and its holding token
}

// Decoder constructor
func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{
		scan: scanner.Scanner{
			Mode: scanner.GoTokens,
		},
	}
	lex.scan.Init(r)
	lex.next() // get next token
	return &Decoder{lex: lex}
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("t")
		} else {
			buf.WriteString("nil")
		}

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem()) // "reflect.Value" of "Interface" or that pointed element

		// ("xxx" "yyy" ...)
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			// "reflect.Value" of each element in that Slice
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	// ((Title "Dr. Strangelove")...)
	case reflect.Struct:
		buf.WriteByte('(')
		// i = 0,1,..,5
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			// Field(i): returns "reflect.Value" of the "value" of that Struct entry
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name) // ((Title
			if err := encode(buf, v.Field(i)); err != nil {  // ((Title "Dr. Strangelove"
				return err
			}
			buf.WriteByte(')') //((Title "Dr. Strangelove")
		}
		// ((Title "Dr. Strangelove") (Subtitle "xxx"), ...)
		buf.WriteByte(')')

	// (("Dr. Strangelove" "Peter Sellers") ("xxx" "yyy") ...)
	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() { // key: reflect.Value of "Dr. Strangelove"
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			// (("Dr. Strangelove"
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			// (("Dr. Strangelove" "Peter Sellers"
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	// all good
	return nil
}

// Marshal: encodes a Go value in S-expression form
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Unmarshal part
// v: reflect.Value of the output Go-object initially; later: reflect.Value of the "value" (2nd) part of a struct's entry
func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	// "nil" or struct field name (but WITHOUT "" or '')
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}

	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // "Dr. Strangelove"
		v.SetString(s)                      // update "v"
		lex.next()                          // ')'
		return

	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i)) // update "v"
		lex.next()         // ')'
		return

	case '(':
		lex.next()       // get the next token, like next '('
		readList(lex, v) // process based on "v"'s type; v is Slice/Map/ etc.
		lex.next()       // consumes ')'
		return
	}
}

// v: reflect.Value of an addressible "Movie" object initially; later: reflect.Value of the "value" (2nd) part of a struct's entry e.g. Slice
func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item1 item2 ...); v has fixed size
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: // (item1 item2 ...)
		// LOOP
		for !endList(lex) {
			// v.Type() is []string => Elem() is "string-type" => New() returns "*string" object => item is addressable "string" object
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name1 value1), (name2 value2), ...)
		// LOOP: reading through each (name, value)
		for !endList(lex) {
			lex.consume('(') // pop out '(' char (of each entry) and move to the next token
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()             // e.g. "Title"/"Year"/"Actor"/"Oscars"
			lex.next()                     // '('
			read(lex, v.FieldByName(name)) // reflect.Value of the "value" (2nd) part of the "Oscars" entry
			lex.consume(')')               // pop out ')' char (of each entry) and move to the next token
		}

	case reflect.Map: // ((key value), (key value), ...)
		v.Set(reflect.MakeMap(v.Type())) // map[string]string
		// LOOP until ')' or EOF
		for !endList(lex) {
			lex.consume('(')                          // move to "Dr. Strangelove"
			key := reflect.New(v.Type().Key()).Elem() // v.Type().Key() is "string-type" => New() returns "*string" => "key" is reflect.Value of a string
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem() // v is map[string]string => v.Type().Elem() return "string-type" => New() returns "*string" => "value" is reflect.Value of a string
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}

}

// has lexer reached EOF yet?
func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

// Decoder main method
func (d *Decoder) Decode(out interface{}) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", d.lex.scan.Position, x)
		}
	}()
	read(d.lex, reflect.ValueOf(out).Elem())
	return nil
}

func main() {
	var strangelove = Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode
	data, err := Marshal(strangelove)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encoded:\n%s\n", string(data))

	// Decode
	dec := NewDecoder(bytes.NewReader(data)) // for streaming: bytes.NewReader(conn)
	var decoded Movie
	err = dec.Decode(&decoded)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nDecoded:\n%+v\n", decoded) // %+v means "print field names and values"

	// // Decode with streaming (not used here)
	// conn, _ = net.Dial(...)
	// dec := NewDecoder(conn)
	// for {
	// 	var m Movie
	// 	if err := dec.Decode(&m); err == io.EOF {
	// 		break
	// 	}
	// 	// do something with "m"
	// }
}
