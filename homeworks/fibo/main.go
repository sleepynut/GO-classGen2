package main

import "fmt"

func main() {
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibo() func() int {
	cur, nxt := 0, 1
	return func() int {
		cur, nxt = nxt, cur+nxt
		return nxt - cur
	}
}
