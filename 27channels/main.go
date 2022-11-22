package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")
	myChan := make(chan int, 5)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	//RECEIVE Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		//fmt.Println(<-myChan)
		val, isChannelopen := <-myChan
		fmt.Println(isChannelopen)
		fmt.Println(val)
		wg.Done()
	}(myChan, wg)
	//SEND Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChan <- 5
		close(myChan)

		//myChan <- 6
		wg.Done()
	}(myChan, wg)

	wg.Wait()
}
