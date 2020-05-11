package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, l := range sites {
		go checkLink(l, c)
	}
	// be sure to avoid referencing the same variable in more than one routine
	// BAD example:
	// for l := range c {
	// 	go func() {
	// 		time.Sleep(1e9)
	// 		checkLink(l, c)
	// 	}()
	// }

	// GOOD example: (GO is pass by value, second copy of var l is created for reference inside lambda)
	for l := range c {
		go func(l string) {
			time.Sleep(1e9)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "is down")
		c <- "might be down"
		return
	}
	fmt.Println(l, "is up")
	c <- l
	return
}
