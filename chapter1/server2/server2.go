// Server returns Path (e.g. "/help"), but also counts the number of requests
// $ go run server2.go & (in Windows no & sign)
// open a new terminal and "./fetch http://localhost:8000" or "./fetch http://localhost:8000/count"
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex // lock the variable "count" when it's being changed
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// shows number of calls to "/" so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
