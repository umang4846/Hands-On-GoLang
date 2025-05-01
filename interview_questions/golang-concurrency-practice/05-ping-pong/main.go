package main

import (
	"fmt"
	"time"
)

func ping(pingCh, pongCh chan bool) {
	for {
		<-pingCh
		fmt.Println("Ping")
		time.Sleep(500 * time.Millisecond)
		pongCh <- true
	}
}

func pong(pingCh, pongCh chan bool) {
	for {
		<-pongCh
		fmt.Println("Pong")
		time.Sleep(500 * time.Millisecond)
		pingCh <- true
	}
}

func main() {
	pingCh := make(chan bool)
	pongCh := make(chan bool)

	go ping(pingCh, pongCh)
	go pong(pingCh, pongCh)

	pingCh <- true // start the game

	time.Sleep(5 * time.Second)
}
