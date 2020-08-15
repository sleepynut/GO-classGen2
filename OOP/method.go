package main

import "fmt"

type rectangle struct {
	width, length int
}

func area(r rectangle) int {
	return r.width * r.length
}

func main() {
	r := rectangle{3, 4}
	fmt.Println(area(r))
}
