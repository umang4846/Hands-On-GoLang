package main

import (
	"fmt"
)

// Function to calculate the sum of elements in a subarray
func sum(arr []int, ch chan<- int) {
	total := 0
	for _, num := range arr {
		total += num
	}
	ch <- total // Sending the sum to the channel
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numWorkers := 4
	chunkSize := len(arr) / numWorkers

	ch := make(chan int, numWorkers)

	// Start multiple goroutines to calculate the sum of subarrays concurrently
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(arr)
		}
		go sum(arr[start:end], ch)
	}

	// Collect results from all goroutines
	total := 0
	for i := 0; i < numWorkers; i++ {
		total += <-ch // Receiving the sum from the channel
	}

	close(ch)

	fmt.Println("Sum of elements:", total)
}
