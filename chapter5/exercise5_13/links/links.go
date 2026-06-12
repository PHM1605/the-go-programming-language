package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// apply <pre> and <post> function for each Node
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	// actual work for THIS Node's children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// Send GET request to <URL>, return links in that URL
func Extract(url string) ([]string, error) {
	// get JSON
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	// convert JSON to Node
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	// start collecting links into <links>
	var links []string
	visitNode := func(n *html.Node) {
		// only work when meet <a href="xxx"> Node
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				// convert relative URL to remote absolute URL
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}

	// we do only <pre> work for this Node
	forEachNode(doc, visitNode, nil)
	return links, nil
}
