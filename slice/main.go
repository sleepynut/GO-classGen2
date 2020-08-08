package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[0:6]
	fmt.Printf("%T\n", primes)
	fmt.Printf("%#v\n", s)
}
