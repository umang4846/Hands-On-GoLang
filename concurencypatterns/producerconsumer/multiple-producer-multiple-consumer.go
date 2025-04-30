package main

import (
	"fmt"
	"sync"
	"time"
)

var messages4 = [][]string{
	{
		"The world itself's",
		"just one big hoax.",
		"Spamming each other with our",
		"running commentary of bullshit,",
	},
	{
		"but with our things, our property, our money.",
		"I'm not saying anything new.",
		"We all know why we do this,",
		"not because Hunger Games",
		"books make us happy,",
	},
	{
		"masquerading as insight, our social media",
		"faking as intimacy.",
		"Or is it that we voted for this?",
		"Not with our rigged elections,",
	},
	{
		"but because we wanna be sedated.",
		"Because it's painful not to pretend,",
		"because we're cowards.",
		"- Elliot Alderson",
		"Mr. Robot",
	},
}

const producerCount4 int = 4
const consumerCount4 int = 3

func multipleProducer4(link chan<- string, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, msg := range messages4[id] {
		link <- msg
	}
}

func multipleConsumer4(link <-chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range link {
		fmt.Printf("Message \"%v\" is consumed by consumer %v\n", msg, id)
		time.Sleep(10 * time.Nanosecond)
	}
}

func main() {
	link := make(chan string)
	wp := sync.WaitGroup{}
	wc := sync.WaitGroup{}

	wp.Add(producerCount4)
	wc.Add(consumerCount4)

	for i := 0; i < producerCount4; i++ {
		go multipleProducer4(link, i, &wp)
	}

	for i := 0; i < consumerCount4; i++ {
		go multipleConsumer4(link, i, &wc)
	}

	wp.Wait()
	close(link)
	wc.Wait()
}
