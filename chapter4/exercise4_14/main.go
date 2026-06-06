package main

import (
	"exercise4_14/github"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	result        *github.IssuesSearchResult
	issuesTpl     = template.Must(template.ParseFiles("templates/issues.html"))
	usersTpl      = template.Must(template.ParseFiles("templates/users.html"))
	userTpl       = template.Must(template.ParseFiles("templates/user.html"))
	milestonesTpl = template.Must(template.ParseFiles("templates/milestones.html"))
	milestoneTpl  = template.Must(template.ParseFiles("templates/milestone.html"))
)

func issuesHandler(w http.ResponseWriter, r *http.Request) {
	issuesTpl.Execute(w, result)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// set of Users (set of authors of all Issues)
	userMap := make(map[string]bool)
	for _, issue := range result.Items {
		userMap[issue.User.Login] = true
	}
	// convert <set> to <string[]>
	var users []string
	for user := range userMap {
		users = append(users, user)
	}
	// set users (string[]) to the users.html template
	usersTpl.Execute(w, users)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// get User's name
	login := strings.TrimPrefix(r.URL.Path, "/user/")
	// collect which Issues is from him
	var issues []*github.Issue
	for _, issue := range result.Items {
		if issue.User.Login == login {
			issues = append(issues, issue)
		}
	}
	// prepare data to be parsed to template
	// 1. define struct 2. use struct as constructor
	data := struct {
		User   string
		Issues []*github.Issue
	}{
		User:   login,
		Issues: issues,
	}
	userTpl.Execute(w, data)
}

func milestonesHandler(w http.ResponseWriter, r *http.Request) {
	// set of milestones from all Issues
	setMilestones := make(map[string]bool)
	for _, issue := range result.Items {
		if issue.Milestone != nil {
			setMilestones[issue.Milestone.Title] = true
		}
	}
	// convert to []string to be parsed to template
	var milestones []string
	for m := range setMilestones {
		milestones = append(milestones, m)
	}
	milestonesTpl.Execute(w, milestones)
}

func milestoneHandler(w http.ResponseWriter, r *http.Request) {
	// URL.Path: http://localhost:8000/milestone/Go1.27
	title := strings.TrimPrefix(r.URL.Path, "/milestone/")
	// list of Issues belonging to that Milestone
	var issues []*github.Issue
	for _, issue := range result.Items {
		if issue.Milestone != nil && issue.Milestone.Title == title {
			issues = append(issues, issue)
		}
	}
	// prepare data (title & list of Issues of a Milestone) to be parsed to template "milestone.html"
	data := struct {
		Title  string
		Issues []*github.Issue
	}{
		Title:  title,
		Issues: issues,
	}
	// parse template
	milestoneTpl.Execute(w, data)
}

func main() {
	// fetch Issues from Github
	var err error
	result, err = github.SearchIssues(
		[]string{"repo:golang/go", "commenter:gopherbot", "json"},
	)
	if err != nil {
		log.Fatal(err)
	}
	// webserver endpoints
	http.HandleFunc("/", issuesHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/user/", userHandler)
	http.HandleFunc("/milestones", milestonesHandler)
	http.HandleFunc("/milestone/", milestoneHandler)

	log.Println("Listening on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
