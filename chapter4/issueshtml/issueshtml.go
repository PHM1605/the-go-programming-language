// template example with HTML
// go run issueshtml.go repo:golang/go commenter:gopherbot json encoder >issues.html
// go run issueshtml.go repo:golang/go 3133 10535 >issues2.html
// (choose 2 issues IDs with special characters like & and <)
package main

import (
	"html/template"
	"issueshtml/github"
	"log"
	"os"
	"time"
)

// calculate "how many days ago" is that time <t>
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var issueList = template.Must(
	template.New("issuelist").
		Parse(`
		<h1>{{.TotalCount}} issues</h1>
		<table>
			<tr style='text-align: left'>
				<th>#</th>
				<th>State</th>
				<th>User</th>
				<th>Title</th>
			</tr>
			{{range .Items}}
			<tr>
				<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
				<td>{{.State}}</td>
				<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
				<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
			</tr>
			{{end}}
		</table>
	`))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// augment the template
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
