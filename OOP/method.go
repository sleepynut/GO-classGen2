package main

import "fmt"

type rectangle struct {
	width, length int
}

func (r *rectangle) area() int {
	return r.width * r.length
}

func main() {
	r := rectangle{3, 4}
	fmt.Println(r.area())
}
