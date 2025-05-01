package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processed %d\n", id, j)
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}
