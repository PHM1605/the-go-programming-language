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
	// NEW
	mux := http.NewServeMux()

	// // NEW: HandlerFunc is a "function type that has method", method ServeHttp() => satisfy Handler interface
	// // type HandlerFunc func(w ResponseWriter, r *Request)
	// // func (f HandlerFunc) ServeHTTP(...)
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))

	// Shorten way of above
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
