package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	oddChan, evenChan := make(chan bool), make(chan bool)

	wg.Add(2)
	go oddPrint(oddChan, evenChan)
	go evenPrint(oddChan, evenChan)

	oddChan <- true
	wg.Wait()
}

func oddPrint(oddChan, evenChan chan bool) {
	defer wg.Done()
	for i := 1; i <= 25; i = i + 2 {
		<-oddChan
		fmt.Println(i)
		if i+1 <= 25 {
			evenChan <- true
		} else {
			close(evenChan)
		}

	}
}

func evenPrint(oddChan, evenChan chan bool) {
	defer wg.Done()
	for i := 2; i <= 25; i = i + 2 {
		_, ok := <-evenChan
		if !ok {
			break
		}
		fmt.Println(i)
		oddChan <- true
	}
}
