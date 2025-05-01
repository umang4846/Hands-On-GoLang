package main

import (
	"fmt"
	"time"
)

func worker(id int, done <-chan struct{}) {
	<-done
	fmt.Printf("Worker %d exiting\n", id)
}

func main() {
	done := make(chan struct{})
	for i := 1; i <= 3; i++ {
		go worker(i, done)
	}

	time.Sleep(1 * time.Second)
	close(done)

	time.Sleep(1 * time.Second)
}
