package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms":      {"data structures"},
	"data structures": {"discrete math"},
	"discrete math":   {"intro to programming", "calculus"},
}

// f: crawl from a course (string) => list of prerequisites ([]string)
// worklist: list of courses to crawl (always 1 here, as key is only a string)
func breadthFirst(f func(string) []string, worklist []string) {
	// mark which course is done added
	seen := make(map[string]bool)
	// old structure of the website crawling program
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				fmt.Println(item)
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// return list of prerequisites of a <course>
func crawl(course string) []string {
	return prereqs[course]
}

func main() {
	// crawl the FULL list of prerequisites for the course "algorithms"
	breadthFirst(crawl, []string{"algorithms"})
}
