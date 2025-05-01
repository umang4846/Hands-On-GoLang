package main

import (
	"fmt"
	"sync"
)

/*
Problem statement : print 1 to 10 in the correct sequence,
with even number printed by one go routine and odd number
printed by another go routine
*/

var wg sync.WaitGroup

func main() {

	oddChan := make(chan bool, 1)
	evenChan := make(chan bool, 1)

	wg.Add(2)
	go printOdd(oddChan, evenChan)
	go printEven(oddChan, evenChan)
	evenChan <- true

	wg.Wait()
}

func printOdd(oddChan, evenChan chan bool) {
	for i := 1; i <= 10; i = i + 2 {
		<-evenChan
		fmt.Println(i)
		oddChan <- true
	}
	wg.Done()
}

func printEven(oddChan, evenChan chan bool) {
	for i := 2; i <= 10; i = i + 2 {
		<-oddChan
		fmt.Println(i)
		evenChan <- true
	}
	wg.Done()
}
