package main

import "fmt"

func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go fibonacci(7, ch)

	for val := range ch {
		fmt.Println(val)
	}
}
