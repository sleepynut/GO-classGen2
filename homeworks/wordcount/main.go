package main

import (
	"fmt"
	"regexp"
	"strings"
)

func WordCount(s string) map[string]int {
	a := make(map[string]int)
	s = regexp.MustCompile("[^a-zA-Z ]").ReplaceAllString(s, "")
	for _, w := range strings.Fields(s) {
		a[w]++
	}
	return a
}
func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s)
	fmt.Println(w)
}
