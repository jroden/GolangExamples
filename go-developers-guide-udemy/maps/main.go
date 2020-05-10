package main

import "fmt"

func main() {
	m := map[string]string{
		"jim":  "red",
		"jane": "blue",
	}
	m["sal"] = "purple"
	delete(m, "jim")
	fmt.Printf("%+v\n", m)
	fmt.Println(m["jim"])
}
