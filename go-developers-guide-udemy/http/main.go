package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// type logWriter implements Write interface
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(resp)

	// create an empty byte slice with 999999 elements--
	// the read function isn't set up to resize slice if it runs out of space
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// use io helper function to write out response body
	// io.Copy(os.Stdout, resp.Body)

	// example using our implementation of write interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// receiver function Write implement Write interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this many bytes:", len(bs))
	return len(bs), nil
}
