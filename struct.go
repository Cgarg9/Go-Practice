package main

import "fmt"

// Define a struct
type person struct {
	name string
	age int
}
// different way of passing argument if it is struct
func (p person) display() {
	fmt.Println("name is", p.name)
	fmt.Println("age is", p.age)
}
func main() {
	a := person{name: "Chirag", age: 21}
	a.display()
}