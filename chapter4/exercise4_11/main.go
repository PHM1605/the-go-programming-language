package main

import (
	"exercise4_11/github"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		usageInfo()
	}
	client, err := github.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	// create/read/update/close
	command := os.Args[1]
	// owner/repo
	parts := strings.Split(os.Args[2], "/")
	if len(parts) != 2 {
		log.Fatal("repo must be owner/repo")
	}
	owner := parts[0]
	repo := parts[1]

	switch command {
	case "create":
		createIssue(client, owner, repo)
	case "read":
		if len(os.Args) < 4 {
			usageInfo()
		}
		readIssue(client, owner, repo, os.Args[3])
	case "update":
		if len(os.Args) < 4 {
			usageInfo()
		}
		updateIssue(client, owner, repo, os.Args[3])
	case "close":
		if len(os.Args) < 4 {
			usageInfo()
		}
		closeIssue(client, owner, repo, os.Args[3])
	default:
		usageInfo()
	}
}

func createIssue(client *github.Client, owner string, repo string) {
	text, err := edit("Issue title\n\nDescribe the issue here...")
	if err != nil {
		log.Fatal(err)
	}

	// start reading content of that "issue-123.md" file
	lines := strings.SplitN(text, "\n", 2) // 2 means at most 2 substrings
	title := strings.TrimSpace(lines[0])
	body := ""
	if len(lines) > 1 {
		body = lines[1]
	}

	// create Issue
	issue, err := client.CreateIssue(owner, repo, title, body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created issue #%d\n%s\n", issue.Number, issue.HTMLURL)
}

func readIssue(client *github.Client, owner, repo, number string) {
	issue, err := client.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("#%d Title: %s\nState: %s\nAuthor: %s\n%s\n\n", issue.Number, issue.Title, issue.State, issue.User.Login, issue.Body)
}

func updateIssue(client *github.Client, owner, repo, number string) {
	issue, err := client.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	initial := issue.Title + "\n\n" + issue.Body
	// open "edit" in Vim, let user fix, and return content-text
	text, err := edit(initial)
	if err != nil {
		log.Fatal(err)
	}
	// process what user types
	lines := strings.SplitN(text, "\n", 2) // 2 = maximum number of substrings; 1 is title and 1 is description
	title := strings.TrimSpace(lines[0])
	body := ""
	if len(lines) == 2 {
		body = lines[1]
	}
	if err := client.UpdateIssue(owner, repo, number, title, body); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Issue updated")
}

func closeIssue(client *github.Client, owner, repo, number string) {
	if err := client.CloseIssue(owner, repo, number); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Issue closed")
}

func edit(initial string) (string, error) {
	// create a file <tmp> in directory "" (current); with name "issue-xyz.md" with <xyz> is random
	tmp, err := os.CreateTemp("", "issue-*.md")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmp.Name())

	// Write text to that <tmp> file
	if _, err := tmp.WriteString(initial); err != nil {
		return "", err
	}
	tmp.Close()

	// Choose a text editor based on user's current settings
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	// like running "vi issue-xyz.md"
	cmd := exec.Command(editor, tmp.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}

	// Read data from "issue-xyz.md" that user has entered
	data, err := os.ReadFile(tmp.Name())
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func usageInfo() {
	log.Fatal(
		"usage\n" +
			"issue create owner/repo\n" +
			"issue read owner/repo number\n" +
			"issue update owner/repo number\n" +
			"issue close owner/repo number",
	)
}
