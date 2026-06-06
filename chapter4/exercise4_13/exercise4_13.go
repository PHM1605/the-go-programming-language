// API to get POSTER IMAGE of a movie at omdbapi.com
// go mod init
// go get github.com/joho/godotenv
// go run . The Matrix
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Movie struct {
	Title  string `json:"Title"`
	Poster string `json:"Poster"`
}

func searchMovie(apiKey, title string) (*Movie, error) {
	myLink := fmt.Sprintf("https://omdbapi.com/?apikey=%s&t=%s", apiKey, url.QueryEscape(title))
	resp, err := http.Get(myLink)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse movie object
	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, err
	}
	// all good
	return &movie, nil
}

func sanitize(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, ":", "_")
	return name
}

func downloadPoster(posterURL, filename string) error {
	resp, err := http.Get(posterURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// parse response to real image
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal(`usage: exercise4_13 "movie title"`)
	}
	// load API key from omdbapi.com
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("OMDB_API_KEY")
	if apiKey == "" {
		log.Fatal("OMDB_API_KEY not found")
	}

	title := strings.Join(os.Args[1:], " ")
	movie, err := searchMovie(apiKey, title)
	if err != nil {
		log.Fatal(err)
	}

	// check to download poster of that movie
	if movie.Poster == "" || movie.Poster == "N/A" {
		log.Fatal("poster not available")
	}
	filename := sanitize(movie.Title) + ".jpg"
	if err := downloadPoster(movie.Poster, filename); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("saved poster to %s\n", filename)
}
