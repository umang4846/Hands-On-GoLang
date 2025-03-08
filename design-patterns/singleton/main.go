package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		//Run the singleton.go file
		//go getInstance()

		//run the singleton.go file
		go getInstance1()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	_, _ = fmt.Scanln()
}
