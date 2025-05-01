package main

import "fmt"

func main() {
	ch := make(chan int)

	// This will cause a deadlock
	ch <- 1
	fmt.Println(<-ch)
}
