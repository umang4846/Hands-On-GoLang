package main

import (
	"fmt"
	"time"
)

func ping(pingCh, pongCh chan bool) {
	for i := 0; i < 5; i++ {
		<-pingCh
		fmt.Println("ping")
		pongCh <- true
	}
}

func pong(pingCh, pongCh chan bool) {
	for i := 0; i < 5; i++ {
		<-pongCh
		fmt.Println("pong")
		pingCh <- true
	}
}

func main() {
	pingCh := make(chan bool)
	pongCh := make(chan bool)

	go ping(pingCh, pongCh)
	go pong(pingCh, pongCh)

	pingCh <- true
	time.Sleep(time.Second)
}
