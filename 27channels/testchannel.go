package main

import (
	"fmt"
	"sync"
)

var wgroup sync.WaitGroup

func main() {
	oddChan := make(chan bool, 2)

	fmt.Println("main ", len(oddChan), cap(oddChan))
	wgroup.Add(2)
	go receive(oddChan)
	go send(oddChan)

	wgroup.Wait()
}
func receive(odd chan bool) {
	defer wgroup.Done()
	fmt.Println("receive ", len(odd), cap(odd))

	odd <- true
	<-odd
}
func send(odd chan bool) {
	defer wgroup.Done()
	fmt.Println("send ", len(odd), cap(odd))
	odd <- true
	fmt.Println("send 2", len(odd), cap(odd))
}
