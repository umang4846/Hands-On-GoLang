package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("Produced:", i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		val := <-ch
		fmt.Println("Consumed:", val)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func main() {
	ch := make(chan int, 2)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}
