package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32

// To avoid concurrency
type database struct {
	mu    sync.Mutex
	items map[string]dollars
}

// NOTE: we must use (*db).list() because we are changing state of <db> inside
func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	err := listTemplate.Execute(w, db.items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// /price?item=shoes
func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%v\n", price)
}

// /create?item=shoes&price=12.3
func (db *database) create(w http.ResponseWriter, req *http.Request) {
	// Get information from query
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	if item == "" {
		http.Error(w, "missing item parameter", http.StatusBadRequest)
		return
	}
	pr, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		http.Error(w, "invalid price", http.StatusBadRequest)
		return
	}
	// Interact db
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, exists := db.items[item]; exists {
		http.Error(w, "item already exists", http.StatusConflict)
		return
	}
	db.items[item] = dollars(pr)
	fmt.Fprintf(w, "created %s at %v\n", item, dollars(pr))
}

// /update?item=socks&price=14
func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceStr := req.URL.Query().Get("price")
	pr, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		http.Error(w, "invalid price", http.StatusBadRequest)
		return
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.items[item]; !exists {
		http.Error(w, "item does not exist", http.StatusNotFound)
		return
	}
	db.items[item] = dollars(pr)
	fmt.Fprintf(w, "updated %s to %v\n", item, dollars(pr))
}

// /delete?item=shoes
func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.items[item]; !exists {
		http.Error(w, "Item does not exist", http.StatusNotFound)
		return
	}

	delete(db.items, item)
	fmt.Fprintf(w, "deleted %s\n", item)
}

// NEW: prepare for the /list endpoint
// "." means map[string]dollars
var listTemplate = template.Must(template.New("list").Parse(`
	<html>
		<head>
			<title>Inventory</title>
		</head>
		<body>
			<h1>Inventory</h1>
			<table border="1">
				<tr>
					<th>Item</th>
					<th>Price</th>
				</tr>
				
				{{range $item, $price := .}}
				<tr>
					<td>{{$item}}</td>
					<td>{{$price}}</td>
				</tr>
				{{end}}
			</table>
		</body>
	</html>
`))

func main() {
	db := &database{
		items: map[string]dollars{
			"shoes": 50,
			"socks": 5,
		},
	}
	// Short way of using "DefaultServeMux"
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // nil = we use DefaultServeMux
}
