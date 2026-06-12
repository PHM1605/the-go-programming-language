// Wait for server ready for 1 minute
// go run main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// with retries for 1 minute
// reports error if all attempts fail
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential backoff
	}
	// eventually failed
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	// url := "https://golang.org"
	// failed server
	url := "http://localhost:9999"
	if err := WaitForServer(url); err != nil {
		// // Method 1
		// fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		// os.Exit(1)

		// Method 2
		// Bonus
		// log.SetPrefix("wait: ")
		// log.SetFlags(0)
		log.Fatalf("Site is down: %v\n", err)

		// Method 3: log the error
		// log.Printf("ping failed: %v; networking disabled", err)
		// Method 4
		// return fmt.Errorf("failed to xxx: %v", err)
	}
	// now server is ready to be used
	fmt.Printf("%s is up!\n", url)
}
