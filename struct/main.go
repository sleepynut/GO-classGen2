package main

import "fmt"

type Vertext struct {
	X int
	Y int
}

type rectangle struct {
	width, length int
}

type Student struct {
	Name, Class string
}

func (s *Student) info() {
	fmt.Printf("Name: %s Class:%s\n", s.Name, s.Class)
}

func (r *rectangle) area() int {
	return r.width * r.length
}

func main() {
	r := rectangle{3, 4}
	fmt.Println(r.area())

	//=======================
	m := map[string]Student{
		"64930495": {Name: "AnuchitO"},
		"333":      {Name: "Panat", Class: "com sci"},
		"555":      {Name: "Adisorn", Class: "xxxx"},
	}

	for _, v := range m {
		v.info()
	}
	//======================

	{
		xn := []int{1, 2, 3, 4, 5}
		fmt.Println("old xn:", xn)

		xnn := xn[1:4:4]
		fmt.Println("sliced xn:", xnn)

		xnn = append(xnn, 99)
		fmt.Println("new xn:", xn)
		fmt.Println("new xnn:", xnn)
	}
}
