package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	chanel := make(chan bool)
	wg.Add(2)
	go oddPrint(chanel)
	go evenPrint(chanel)
	wg.Wait()
}
func oddPrint(chanel chan bool) {
	defer wg.Done()
	for i := 1; i <= 25; i = i + 2 {
		fmt.Println(i)
		chanel <- true
		<-chanel

	}
}
func evenPrint(chanel chan bool) {
	defer wg.Done()
	for i := 2; i <= 25; i = i + 2 {
		<-chanel
		fmt.Println(i)
		chanel <- true
	}
}
