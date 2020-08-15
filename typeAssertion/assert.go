package main

import "fmt"

type Samsung struct {
	version string
}

func main() {
	var i interface{} = []string{"MY", "GOD"}
	i = &Samsung{version: "s10+"}
	// s := i.(string)
	// fmt.Println(s)
	switch i.(type) {
	case nil, int, string, []int:
		fmt.Printf("%T\n", i)
	case *Samsung, Samsung:
		p, _ := i.(*Samsung)
		fmt.Printf("Version of my phone %s\n", p.version)
	default:
		fmt.Printf("%T\n", i)
	}
}
