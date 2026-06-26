package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type job struct {
	url string
}

var (
	rootHost string

	worklist   = make(chan []job) // can be duplicated
	unseenJobs = make(chan job)   // de-duplicate links

	seen = make(map[string]bool) // to make sure we aren't stucked in a circular (linkA contains linkB, linkB also contains linkA)
	mu   sync.Mutex              // for carefully modify "seen" (set of links)
)

func localFileName(u *url.URL) string {
	path := u.Path // "" or "/" for root; "customer" for deeper
	if path == "" || strings.HasSuffix(path, "/") {
		path += "index.html" // "index.html" eventually for root
	}
	// if path is https://golang.org/customer
	if filepath.Ext(path) == "" {
		path += ".html" // "customer/customer.html"
	}
	return filepath.Join("mirror", u.Host, path) // mirror/golang.org/index.html for root; mirror/golang.org/customer/customer.html for deeper pages
}

func localPath(u *url.URL) string {
	path := u.Path // "/" or "/customer"
	// /index.html for root
	if path == "" || strings.HasSuffix(path, "/") {
		path += "index.html"
	}
	// /customer/customer.html for deeper page
	if filepath.Ext(path) == "" {
		path += ".html"
	}
	return path
}

// create html folder local; return links inside that pageURL in the meantime
func mirrorPage(pageURL string) ([]string, error) {
	// get HTML for that page
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	// convert HTML string to Node for traversing later
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	// convert string to e.g. *URL contains "golang.org"
	base, _ := url.Parse(pageURL)

	// collect links to be collected here
	var links []string
	visitNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}
				// convert href="/doc/" inside to "http://golang.org/doc/"
				link, err := base.Parse(attr.Val) // return *URL for "http://golang.org/doc/"
				if err != nil {
					continue
				}
				// only consider same domain
				if link.Host != rootHost {
					continue
				}
				link.Fragment = ""   // http://golang.org/doc/#a and http://golang.org/doc/#b are same link
				abs := link.String() // convert to string
				links = append(links, abs)
				// change to href="/customer/customer.html"
				n.Attr[i].Val = localPath(link)
			}
		}
	})
	// creating that .html file in local
	filename := localFileName(base)                 // /mirror/golang.org/index.html
	err = os.MkdirAll(filepath.Dir(filename), 0755) // make /mirror/golang.org parent folder
	if err != nil {
		return nil, err
	}
	file, err := os.Create(filename) // index.html
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// put content into that file
	err = html.Render(file, doc)
	if err != nil {
		return nil, err
	}

	return links, nil
}

func visitNode(n *html.Node, pre func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	// dig deeper to do "pre" for children nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visitNode(c, pre)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: ./exercise8_7 <url")
	}
	root := os.Args[1] // https://golang.org
	// convert string to Node
	u, err := url.Parse(root) // https://golang.org
	if err != nil {
		log.Fatal(err)
	}
	rootHost = u.Host // golang.org

	// add link (root) from command line first
	go func() {
		worklist <- []job{{url: root}}
	}()

	// start 20 crawlers
	for i := 0; i < 20; i++ {
		go func() {
			for j := range unseenJobs {
				// for this crawling we print the link being fetched
				fmt.Println("fetch: ", j.url)
				links, err := mirrorPage(j.url)
				if err != nil {
					log.Println(err)
					continue
				}
				// prepare "jobs" from "links"
				var jobs []job
				for _, link := range links {
					jobs = append(jobs, job{url: link})
				}
				// send jobs to "worklist" (can be duplicated); as a separate goroutine; to avoid deadlock
				go func() {
					worklist <- jobs
				}()
			}
		}()
	}
	// de-duplicate links before sending to workers' channel
	for list := range worklist {
		for _, j := range list {
			mu.Lock()
			if !seen[j.url] {
				seen[j.url] = true
				mu.Unlock()
				unseenJobs <- j
			} else {
				mu.Unlock()
			}
		}
	}
}
