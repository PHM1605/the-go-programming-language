package main

import "fmt"

func MyJoin(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for _, s := range strs[1:] {
		result += sep + s
	}
	return result
}

func main() {
	fmt.Println(MyJoin(",", "a", "b", "c"))
	fmt.Println(MyJoin("-", "Go", "Java", "Python"))
}
