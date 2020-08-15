package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	switch i.(type) {
	case nil:
		fmt.Printf("%T\n", i)
	case int:
		fmt.Printf("%T\n", i)
	case string:
		fmt.Printf("%T\n", i)
	case []int:
		fmt.Printf("%T\n", i)
	default:
		fmt.Printf("%T\n", i)
	}
}
