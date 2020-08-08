package main

import "fmt"

type Vertext struct {
	X int
	Y int
}

func main() {
	v := Vertext{1, 2}
	fmt.Printf("%#v\n", v)

	v.X = 15
	v.Y = 20

	fmt.Printf("%+v\n", v)
}
