package main

type messaging interface {
	getGreeting() string
}

type person struct {
	name string
	age  int
}

type robot struct {
	modelno string
}

func main() {
	s := person{name: "sally", age: 3}
	intro(s)
	l := robot{modelno: "2345232442"}
	intro(l)
}

func intro(m messaging) {
	println("hi I am known as", m.getGreeting())
}

func (p person) getGreeting() string {
	return p.name
}

func (p robot) getGreeting() string {
	return p.modelno
}
