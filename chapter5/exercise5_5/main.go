package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	// count words in TextNode
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}
	// count images
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	// recursive
	w1, i1 := countWordsAndImages(n.FirstChild)
	w2, i2 := countWordsAndImages(n.NextSibling)

	words += w1 + w2
	images += i1 + i2

	return words, images
}

// "bare return": return only statement
func CountWordsAndImages(url string) (words, images int, err error) {
	// get HTML from website
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	// convert JSON to Node
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	// count words and images from Node
	words, images = countWordsAndImages(doc)
	return // equals "return words, images, err"
}
func main() {
	words, images, err := CountWordsAndImages("https://golang.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("words: %d\n", words)
	fmt.Printf("images: %d\n", images)
}
