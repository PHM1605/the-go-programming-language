package main

import (
	"exercise5_13/links"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

var rootHost string

// 1st parameter: a function takes a "string" and returns a "[]string" => crawl LINKS from an URL
// 2nd parameter: []string; list of websites
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool) // make which websites have been done

	// 1st iteration: crawl to get a list of URLs from "https://golang.org" and "https://google.com" => [https://golang.org/users, https://google.com/templates]
	// 2nd iteration: dig deeper => [https://golang.org/users/1, https://golang.org/users/abc, https://google.com/templates/xyz]
	// 3rd iteration: crawl(item) finds nothing => escape for loop
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		// iterate over list of websites
		for _, item := range items {
			// convert from string => URL
			u, err := url.Parse(item)
			if err != nil {
				continue
			}
			// take Host from URL to skip
			if u.Host != rootHost {
				continue
			}
			// crawl deep inside
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// save to "mirror/" folder
func savePage(rawURL string) error {
	// convert string to URL
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	// skip foreign domains
	if u.Host != rootHost {
		return nil
	}
	// get content in that URL
	resp, err := http.Get(rawURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// mirror/golang.org
	path := filepath.Join("mirror", u.Host)
	// full link: golang.org or golang.org/
	if u.Path == "" || u.Path == "/" {
		path = filepath.Join(path, "index.html")
	} else { // full link: golang.org/users or golang.org/users/
		path = filepath.Join(path, u.Path)
		// if it's golang.org/users/index.html we don't create manually
		if filepath.Ext(u.Path) == "" {
			path = filepath.Join(path, "index.html")
		}
	}

	// create local dir mirror/golang.org/users/
	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return nil
	}
	// create file *.html inside that dir
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

// crawling function: takes an URL an return list of links
func crawl(rawURL string) []string {
	// print that URL first
	fmt.Println(rawURL)
	// save that URL
	if err := savePage(rawURL); err != nil {
		log.Fatal(err)
	}
	// extract links from that URL
	list, err := links.Extract(rawURL)

	if err != nil {
		log.Print(err)
	}
	return list // list of links from that URL
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s url\n", os.Args[0])
		os.Exit(1)
	}

	u, err := url.Parse(os.Args[1]) // e.g. https://golang.org
	if err != nil {
		log.Fatal(err)
	}
	rootHost = u.Host // golang.org

	breadthFirst(crawl, []string{os.Args[1]}) // ["https://golang.org"]
}
