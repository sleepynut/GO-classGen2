package main

import "fmt"

type student struct {
	name string
	id   int
}

func (s *student) String() string {
	return fmt.Sprintf("[name: %s,id: %d]\n",
		s.name, s.id)
}

func main() {
	s := student{"Nut", 123}
	fmt.Println(&s)
}
