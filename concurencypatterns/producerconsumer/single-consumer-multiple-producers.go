package main

import (
	"fmt"
	"sync"
)

var messages3 = [][]string{
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

const producerCount int = 4

func multipleProducer3(jobs chan<- string, idx int, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, msg := range messages3[idx] {
		fmt.Printf("Producer %v sending message \"%v\"\n", idx, msg)
		jobs <- msg
	}

}

func singleConsumer3(jobs <-chan string, done chan<- bool) {
	for msg := range jobs {
		fmt.Printf("Consumed message \"%v\"\n", msg)
	}
	done <- true
}

func main() {
	jobs := make(chan string)
	done := make(chan bool)

	wg := sync.WaitGroup{}
	for i := 0; i < producerCount; i++ {
		wg.Add(1)
		go multipleProducer3(jobs, i, &wg)
	}

	go singleConsumer3(jobs, done)

	wg.Wait()
	close(jobs)

	<-done

}
