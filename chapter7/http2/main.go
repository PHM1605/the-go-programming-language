package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

// to format what displays when print a "dollars" type variable
func (d dollars) String() string { return fmt.Sprintf("%.2f", d) }

// Special in Golang: Database ATTACHES handler functions to it, to provide data to all endpoints
type database map[string]dollars

// to comply to Handler interface
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// NEW: provide some routes
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}

	case "/price":
		item := req.URL.Query().Get("item") // /price?item=shoes
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func main() {
	// map product and price
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
