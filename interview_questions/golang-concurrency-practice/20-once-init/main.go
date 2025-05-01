package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initOnce() {
	fmt.Println("init called")
}

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			once.Do(initOnce)
		}()
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		once.Do(initOnce)
	}()
	wg.Wait()
}
