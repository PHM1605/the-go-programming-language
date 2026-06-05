// Write a program that prints SHA256 hash of its stdin by default; but support a command-line flag to print SHA384 and SHA512 instead
// -n means "no new line"
// echo -n "hello" | go run exercise4_2.go
// echo -n "hello" | go run exercise4_2.go -algo sha384
// echo -n "hello" | go run exercise4_2.go -algo sha512
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// define flag for command-line arguments => key, default-value, description
	algo := flag.String("algo", "sha256", "hash algorithm: sha256, sha384, sha512")
	flag.Parse()
	data, err := io.ReadAll(os.Stdin) // read the "echo -n hello" part
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading stdin: %v\n", err)
		os.Exit(1)
	}

	switch *algo {
	case "sha256":
		hash := sha256.Sum256(data)
		fmt.Printf("%x\n", hash)
	case "sha384":
		hash := sha512.Sum384(data)
		fmt.Printf("%x\n", hash)
	case "sha512":
		hash := sha512.Sum512(data)
		fmt.Printf("%x\n", hash)
	default:
		fmt.Fprintf(os.Stderr, "unknown algorithm %q, chose: sha256, sha384, sha512\n", *algo)
		os.Exit(1)
	}

}
