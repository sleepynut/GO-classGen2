package main

import "fmt"

type person struct {
	name string
}

func (p *person) walk() {
	fmt.Printf("%s is walking.\n", p.name)
}
func (p *person) eat() {
	fmt.Printf("%s is eating.\n", p.name)
}
func (p *person) greeting() {
	fmt.Printf("Hello! %s\n", p.name)
}
func (p *person) myname() string  { return p.name }
func (p *person) setter(n string) { p.name = n }

func main() {
	p := person{name: "NUT"}
	p.walk()
	p.eat()
	p.greeting()
	fmt.Println("GETTING: ", p.myname())

	p.setter("SOMEONE")
	fmt.Println("GETTING: ", p.myname())
}
