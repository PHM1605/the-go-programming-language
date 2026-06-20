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

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	// map product and price
	db := database{"shoes": 50, "socks": 5}

	// Shorten way of above
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	// NEW: here we choose default "DefaultServeMux" by passing "nil"
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
