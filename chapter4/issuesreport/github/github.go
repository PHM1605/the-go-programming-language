// send request to an API of Github to see "Github issue tracker"
// will be used by "issues.go"
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"` // how many Issues
	Items      []*Issue // pointers to those Issues
}

// make HTTP request and decode result as JSON
// terms: ["repo:golang/go", "is:open", "json", "decoder"]
// => QueryEscape("repo:golang/go is:open json decoder")
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// close the response from Github properly
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	// Decode JSON of response to Go data struct
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	// success
	resp.Body.Close()
	return &result, nil
}
