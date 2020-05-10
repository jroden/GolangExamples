package main

import (
	"fmt"
	"io"
	"os"
)

// assignment is to create a program that reads the contents
// of a text file then prints its contents to the terminal
func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lenwritten, err := io.Copy(os.Stdout, f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("num of bytes written out", lenwritten)
}
