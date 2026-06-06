package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const IssuesURL = "https://api.github.com/search/issues"

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

type Milestone struct {
	Title string `json:"title"`
}

type Issue struct {
	Number    int        `json:"number"`
	HTMLURL   string     `json:"html_url"`
	Title     string     `json:"title"`
	State     string     `json:"state"`
	User      *User      `json:"user"`
	Milestone *Milestone `json:"milestone"`
}

type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue `json:"items"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	// parse to result
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	// all good
	return &result, nil
}
