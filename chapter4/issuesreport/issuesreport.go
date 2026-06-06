// template example with text/template library
// go run issuesreport.go repo:golang/go commenter:gopherbot json decoder
package main

import (
	"issuesreport/github"
	"log"
	"os"
	"text/template"
	"time"
)

// calculate "how many days ago" is that time <t>
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// "." in ".TotalCount" or ".Items" is "github.IssuesSearchResult"
// "." in ".Number", ".User" etc. is pointer-to-Issue (in "Items")
// ".Title" is the argument of "print xxx"
// ".CreatedAt" is the argument for function "daysAgo()"
const templ = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------
Number:	{{.Number}}
User:	{{.User.Login}}
Title: 	{{.Title | printf "%.64s"}}
Age:	{{.CreatedAt | daysAgo}}
{{end}}`

var report = template.Must(
	template.New("issuelist"). // template name
					Funcs(template.FuncMap{"daysAgo": daysAgo}). // pass in which functions to use in template
					Parse(templ),                                // template definition
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// augment the template
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
