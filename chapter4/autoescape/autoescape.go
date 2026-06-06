// template using "template.HTML" instead of "string"
// go run autoescape.go >autoescape.html
package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	// we will see these 2 rendered differently on template
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(
		template.New("escape").Parse(templ),
	)
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
