package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(2 * time.Second):
		fmt.Println("Operation timed out")
	}
}
