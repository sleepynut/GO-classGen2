package main

import "fmt"

type Student struct {
	Name, Class string
}

var m = map[int]Student{
	123: {"Nut", "KBTG"},
	456: {"X", "Somewhere"},
}

func main() {
	keys := []int{123, 456, 789, 000}
	for _, k := range keys {
		m[k] = Student{"X", "YY"}
	}

	for i, k := range m {
		fmt.Println(i, k)
	}
}
