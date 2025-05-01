package main

import "fmt"

func main() {
	buffered := make(chan int, 1)
	unbuffered := make(chan int)

	fmt.Println("Sending to buffered channel")
	buffered <- 10
	fmt.Println("Sent to buffered channel")

	go func() {
		fmt.Println("Sending to unbuffered channel")
		unbuffered <- 20
		fmt.Println("Sent to unbuffered channel")
	}()

	val := <-unbuffered
	fmt.Println("Received from unbuffered channel:", val)
}
