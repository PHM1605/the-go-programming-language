package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Package struct {
	ImportPath string
	Deps       []string
}

// split "fmt\nstrconv\n"
func splitLines(s string) []string {
	var lines []string
	// start: marks start of newline
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			if start < i {
				lines = append(lines, s[start:i])
			}
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

func main() {
	// usage: "./exercise10_4 strconv"
	if len(os.Args) < 2 {
		log.Fatal("usage: ./exercise10_4 package1 package2...")
	}

	// first "go list strconv xxx etc." command => return the import path of "strconv" (or more packages)
	cmd := exec.Command("go", append([]string{"list"}, os.Args[1:]...)...)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	// split "fmt\nstrconv\n" => set ["fmt", "strconv"]
	targets := make(map[string]bool)
	for _, p := range splitLines(string(out)) {
		targets[p] = true
	}

	// 2nd "go list -json all" to list all packages in current workspace
	// => output: {"ImportPath":"fmt",...}, {"ImportPath":"errors",...}, {"ImportPath":"strconv",...}
	cmd = exec.Command("go", "list", "-json", "all")
	stdout, err := cmd.StdoutPipe() // "stdout" is a "Pipe" of above dicts
	if err != nil {
		log.Fatal(err)
	}
	// start "Pipe"
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(stdout) // decode a Pipe

	for {
		var pkg Package
		// extract each dict above to "pkg"
		// {"ImportPath":"fmt", "Deps":[xxx]}
		if err := dec.Decode(&pkg); err != nil {
			break
		}
		for _, dep := range pkg.Deps {
			// NOTE: if list of dependencies of a package has "strconv" or "fmt", mark it
			if targets[dep] {
				fmt.Println(pkg.ImportPath)
				break
			}
		}
	}

	// normal cleanup after cmd.Start()
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
