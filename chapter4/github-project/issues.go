// send request to an API of Github to see "Github issue tracker"
// go run issues.go repo:golang/go is:open json decoder
package main

import (
	"fmt"
	"github-project/github"
	"log"
	"os"
)

func main() {
	// result: list of Issues and number-of-Issues
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		// "-" means "left-align"; "5" means "at-least-5-chars-space"
		// ".55" means "maximum 55 chars"
		// "9" at the beginning means "minimum 9 chars"
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
