package main

import (
	"container/heap"
	"fmt"
)

// Item represents an element in the priority queue
type Item struct {
	priority int    // Sorting criteria (integer)
	value    string // Actual stored data (string)
}

// PriorityQueue implements heap.Interface
type PriorityQueue []Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority > pq[j].priority } // Max Heap
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	pq := &PriorityQueue{
		{priority: 3, value: "http://google.com/"},
		{priority: 1, value: "Very Low"},
		{priority: 5, value: "High"},
		{priority: 4, value: "Medium"},
	}
	heap.Init(pq)

	heap.Push(pq, Item{priority: 6, value: "Very High"}) // Insert new element

	fmt.Println(heap.Pop(pq)) // {6 "Very High"}
	fmt.Println(heap.Pop(pq)) // {5 "High"}
	fmt.Println(heap.Pop(pq)) // {4 "Medium"}
	fmt.Println(heap.Pop(pq)) // {3 "Low"}
	fmt.Println(heap.Pop(pq)) // {1 "Very Low"}
}
