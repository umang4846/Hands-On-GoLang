package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu  sync.Mutex
	val int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter.val)
}
