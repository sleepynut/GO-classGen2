package main

import "fmt"

func main() {
	var v interface{}
	v = 1
	fmt.Printf("%[1]T %#[1]v\n", v)

	v = "1"
	fmt.Printf("%[1]T %#[1]v\n", v)

	v = []int{1, 2, 3, 4, 5}
	fmt.Printf("%[1]T %#[1]v\n", v)

	v = map[string]int{"test": 1, "sample": 2}
	fmt.Printf("%[1]T %#[1]v\n", v)

}
