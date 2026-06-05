// send request to an API of Github to see "Github issue tracker"
// NEW: we will categorize Issues
// - less than 1 month old
// - less than 1 year old
// - more than 1 year old
// go run ./main.go repo:golang/go is:open json decoder
package main

import (
	"exercise4_10/github"
	"fmt"
	"log"
	"os"
	"time"
)

func printGroup(title string, issues []*github.Issue) {
	fmt.Printf("\n%s (%d issues)\n", title, len(issues))
	for _, item := range issues {
		// "-" means "left-align"; "8" means "at-least-8-chars-space"
		// ".155" means "maximum 155 chars"
		// "9" at the beginning means "minimum 9 chars"
		fmt.Printf(
			"#%-8d %10.10s %.155s\n",
			item.Number,
			item.User.Login,
			item.Title,
		)
	}
}

func main() {
	// result: list of Issues and number-of-Issues
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var monthOld []*github.Issue
	var yearOld []*github.Issue
	var older []*github.Issue

	now := time.Now()

	for _, item := range result.Items {
		age := now.Sub(item.CreatedAt)
		switch {
		case age < 30*24*time.Hour:
			monthOld = append(monthOld, item)
		case age < 365*24*time.Hour:
			yearOld = append(yearOld, item)
		default:
			older = append(older, item)
		}
	}

	printGroup("Less than a month old", monthOld)
	printGroup("Less than a year old", yearOld)
	printGroup("More than a year old", older)
}
