package main

import (
	"fmt"
	"sync"
	"time"
)

//task definition
type Task struct {
	ID int
}

// way to process the tasks
func (t *Task) Process() {
	fmt.Println("Processing task %d\n", t.ID)

	time.Sleep(2 * time.Second)
}

//worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

//functions to execute worker pool
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	//initialize the tasks channel
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	//starts workers
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	//send tasks to the tasks channel
	wp.wg.Add(len(wp.Tasks))

	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	//wait for all tasks to finish
	wp.wg.Wait()
}
