package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	a := make(map[string]int)
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
