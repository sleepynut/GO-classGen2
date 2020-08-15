package main

import "fmt"

type Samsung struct {
	version string
}

func (s *Samsung) Info() string {
	return fmt.Sprintf("info: %s", s.version)
}

type OnePlus struct {
	Samsung
	version string
}

// func (op *OnePlus) Info() string {
// 	return fmt.Sprintf("info: %s", op.version)
// }

func main() {
	s := Samsung{}
	s.version = "s10+"
	fmt.Println(s.Info())

	op := OnePlus{}
	op.version = "1+"
	fmt.Println(op.Info())
}
