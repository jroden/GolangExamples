package main

import "fmt"

type shape interface {
	getArea() int
}

type triangle struct {
	length int
}

type square struct {
	length int
}

func main() {
	s := square{length: 3}
	printarea(s)
	t := triangle{length: 3}
	printarea(t)

}

func (t triangle) getArea() int {
	// for this example we will treat length as area
	return t.length
}

func (s square) getArea() int {
	return s.length * s.length
}

func printarea(s shape) {
	fmt.Println("area of shape is", s.getArea())
}
