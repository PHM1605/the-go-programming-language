package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const BaseURL = "https://api.github.com"

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int       `json:"number"`
	HTMLURL   string    `json:"html_url"`
	Title     string    `json:"title"`
	State     string    `json:"state"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
}

type IssueRequest struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
	State string `json:"state,omitempty"`
}

type Client struct {
	Token string
}

func NewClient() (*Client, error) {
	// load file .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	// get value from file .env
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN not set")
	}
	return &Client{Token: token}, nil
}

// Methods of <Client>
func (c *Client) CreateIssue(owner, repo, title, body string) (*Issue, error) {
	// convert object to JSON to be sent to Github
	reqBody, err := json.Marshal(
		IssueRequest{Title: title, Body: body},
	)
	if err != nil {
		return nil, err
	}
	// send POST request
	url := fmt.Sprintf("%s/repos/%s/%s/issues", BaseURL, owner, repo)
	resp, err := c.do(http.MethodPost, url, reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func (c *Client) UpdateIssue(owner, repo, number, title, body string) error {
	// convert from newly created Go object (with updated data) to JSON
	reqBody, err := json.Marshal(
		IssueRequest{Title: title, Body: body},
	)
	if err != nil {
		return err
	}

	// start sending PATCH request
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", BaseURL, owner, repo, number)
	resp, err := c.do(http.MethodPatch, url, reqBody)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil // no error
}

func (c *Client) GetIssue(owner, repo, number string) (*Issue, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", BaseURL, owner, repo, number)
	// Send GET request to Github
	resp, err := c.do(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// parse response to Issue object
	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	// okayish
	return &issue, nil
}

func (c *Client) CloseIssue(owner, repo, number string) error {
	// convert Go object to JSON
	reqBody, err := json.Marshal(
		IssueRequest{State: "closed"},
	)
	if err != nil {
		return err
	}
	// send PATCH request with only "State"="closed"
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", BaseURL, owner, repo, number)
	resp, err := c.do(http.MethodPatch, url, reqBody)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

// create and send request to Github
func (c *Client) do(method, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	// set headers
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")
	// send
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// if 101 or 404 error
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		return nil, fmt.Errorf("github api error: %s", resp.Status)
	}
	// okayish
	return resp, nil
}
