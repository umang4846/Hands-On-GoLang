package main

import "fmt"

func sum(arr []int, ch chan int) {
	total := 0
	for _, val := range arr {
		total += val
	}
	ch <- total
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	ch := make(chan int)

	go sum(nums[:len(nums)/2], ch)
	go sum(nums[len(nums)/2:], ch)

	x, y := <-ch, <-ch
	fmt.Println("Total Sum:", x+y)
}
