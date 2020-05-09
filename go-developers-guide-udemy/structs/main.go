package main

import "fmt"

type contact struct {
	number int
}

type person struct {
	name string
	age  int
	c    contact
}

func main() {
	// p := person{"john", 3}
	// var p person
	p := person{name: "john", age: 3, c: contact{3}} // if multi-line, need commas even on last assignment
	fmt.Printf("%+v", p)
	p.name = "johny"
	fmt.Println("name: ", p.name)

	// pointer example
	ppointer := &p
	ppointer.updateName("jim")
	fmt.Printf("%+v", p)

	// pointer shortcut
	p.updateName("Bob")
	fmt.Printf("%+v", p)

}

// if persistent changes are desired, can use func return or pointers
// Go is pass by value language.  Anytime you pass data to function a copy is made
// Several types of data structures are reference types-- for these the pass by value still refers
//		the original memory location (can update without requiring pointers)
// value types: int, float, string, bool, structs
// reference types: slices, maps, channels, pointers, functions
func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).name = name
}
