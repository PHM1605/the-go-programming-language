// JSON example
// go run movie.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// NOTE: field names must be "capital-first" so that JSON-ify successful
// field-tag: `json:"xxx"`control behavior of "encoding:json" package
// - "released" or "color": alternative names for that field
// - "omitempty": if that field is "false" or "0" then don't include in final JSON
type Movie struct {
	Title  string
	Year   int  `json:"released"`        // last field: field tag
	Color  bool `json:"color,omitempty"` // color movie or black/white movie
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	// convert Go data to JSON => "Marshal" process
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// convert Go data to JSON with indentation
	data2, err2 := json.MarshalIndent(movies, "", "    ") // 2nd: prefix of each line; 3rd: indentation space
	if err2 != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)

	// convert JSON to Go data structure => "Unmarshal"
	var titles []struct{ Title string } // take only "Title" field of data
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // list of {xxx}
}
