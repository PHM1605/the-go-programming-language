package links

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// apply <pre> and <post> functions for each Node
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	// dig deeper recursively
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// Send GET request to <URL>, return links in that URL
// NEW: use "context" for cancelling this goroutine (which calls this Extract()) if needed
func Extract(ctx context.Context, url string) ([]string, error) {
	// NEW: http request that has "context"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	// send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	// convert JSON to Node
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// start collecting links into []links
	var links []string
	visitNode := func(n *html.Node) {
		// only stop when we meet <a href="xxx"> Node
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				// href here; convert relative URL to absolute URL
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	// recursive visit Node, starting from root
	forEachNode(doc, visitNode, nil)
	return links, nil
}
