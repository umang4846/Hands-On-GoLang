package main

import (
	"fmt"
)

func genNums(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}

func genChars(ch chan<- string) {
	for _, c := range []string{"a", "b", "c"} {
		ch <- c
	}
	close(ch)
}

func fanIn(ch1 <-chan int, ch2 <-chan string, out chan<- any) {
	for ch1 != nil || ch2 != nil {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				out <- v
			}
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				out <- v
			}
		}
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	out := make(chan any)

	go genNums(ch1)
	go genChars(ch2)
	go fanIn(ch1, ch2, out)

	for v := range out {
		fmt.Println(v)
	}
}
