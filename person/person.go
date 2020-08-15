package main

import "fmt"

type person struct {
	name    string
	friends map[string]int
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
func (p *person) Name() string { return p.name }
func (p *person) setName(n string) {
	p.name = n
	p.friends["NUT"]++
}

func (p *person) String() string {
	return fmt.Sprintf("[Name: %s, Friend(s): %v]\n", p.Name(), p.friends)
}

func main() {
	p := person{name: "NUT", friends: make(map[string]int)}
	p.walk()
	p.eat()
	p.greeting()
	fmt.Println("GETTING: ", p.Name())

	p.setName("SOMEONE")
	fmt.Println("GETTING: ", p.Name())

	fmt.Println(&p)
}
