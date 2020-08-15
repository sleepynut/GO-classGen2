package main

import (
	"errors"
	"fmt"
)

type BusinessError struct {
	err string
}

func (b *BusinessError) Error() string {
	return b.err
}

func PrintErr(err error) {
	fmt.Println(err)
}
func main() {
	err := errors.New("my error msg")
	PrintErr(err)
	berr := BusinessError{"Business error"}
	PrintErr(&berr)
}
