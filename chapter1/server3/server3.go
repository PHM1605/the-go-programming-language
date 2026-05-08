// Server reports header and form data
// $ go run server3.go & (in Windows no & sign)
// open a new terminal and "./fetch http://localhost:8000"
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// print "Method" "URL" "Prototype"
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// print Header
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	// print Host
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	// print PublicIP:Port
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	// Parsing form data, check if there is error
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	// Print form
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
