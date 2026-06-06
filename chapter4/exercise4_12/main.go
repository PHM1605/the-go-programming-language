package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const CacheFile = "comics.json"

type Comic struct {
	Num        int    `json:"num"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func latestComicNumber() (int, error) {
	// get response
	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	// get comic from response
	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return 0, err
	}
	// all good
	return comic.Num, nil
}

// fetch comic number "123" into pointer-to-Comic
func fetchComic(num int) (*Comic, error) {
	// fetch JSON
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// if error
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("comic %d not found", num)
	}
	// convert JSON to object <Comic>
	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, err
	}
	// all good
	return &comic, nil
}

func downloadAll() ([]Comic, error) {
	latest, err := latestComicNumber()
	if err != nil {
		return nil, err
	}
	fmt.Printf("latest comic: %d\n", latest)
	// list of downloaded comics
	comics := make([]Comic, 0, latest) // 2nd=length, 3rd=capacity
	for i := 1; i <= latest; i++ {
		// fetch comic <i>'s information
		comic, err := fetchComic(i)
		if err != nil {
			continue
		}
		comics = append(comics, *comic)
		// show some progress printing to track
		if i%100 == 0 {
			fmt.Printf("downloaded %d\n", i)
		}
	}
	// convert from Go object to JSON
	data, err := json.MarshalIndent(comics, "", "  ") // 2nd=prefix; 3rd=indent-chars
	if err != nil {
		return nil, err
	}
	// write JSON to cache
	// permission=0644 => owner read/write; group read only; others read only
	if err := os.WriteFile(CacheFile, data, 0644); err != nil {
		return nil, err
	}
	// print to verbose
	fmt.Printf("saved %d comics to %s\n", len(comics), CacheFile)
	// all good
	return comics, nil
}

func loadOrDownload() ([]Comic, error) {
	// os.Stat(): show statistics of that file
	// if no file => err has value => not run inside loop
	if _, err := os.Stat(CacheFile); err == nil {
		fmt.Println("loading local cache...")
		data, err := os.ReadFile(CacheFile)
		if err != nil {
			return nil, err
		}
		// turn JSON cache into meaningful objects
		var comics []Comic
		if err := json.Unmarshal(data, &comics); err != nil {
			return nil, err
		}
		// read JSON cache well
		return comics, nil
	}
	// cache not exist => we download into "json"
	fmt.Println("downloading xkcd archive...")
	return downloadAll()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: xkcd <search-term>")
	}
	comics, err := loadOrDownload()
	if err != nil {
		log.Fatal(err)
	}
	term := strings.ToLower(os.Args[1])
	found := false
	// searching
	for _, comic := range comics {
		text := strings.ToLower(comic.Title + " " + comic.Transcript)
		if strings.Contains(text, term) {
			found = true
			fmt.Printf("https://xkcd.com/%d/\n%s\n%s\n", comic.Num, comic.Title, comic.Transcript)
		}
	}
	if !found {
		fmt.Println("no matches")
	}
}
