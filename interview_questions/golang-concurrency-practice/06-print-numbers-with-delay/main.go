package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumber(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(n)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go printNumber(i, &wg)
	}
	wg.Wait()
}
