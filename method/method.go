package main

import "fmt"

type Day int

func (d Day) Today() string {
	return fmt.Sprintf("today : %d\n", d)
}

func main() {
	var d Day
	d = 2
	today := d.Today()
	fmt.Println(today)
}
