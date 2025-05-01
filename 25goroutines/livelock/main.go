package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
Problem statement
Consider a classic livelock scenario in which a husband and wife named Alice and Bob attempt to eat soup but only have one spoon between them.
Each spouse is overly courteous and will pass the spoon to the other if the other hasn’t eaten yet.

Use the steps below to create this livelock scenario in Golang:
 -> Create a function that creates a lock on resource x (a spoon in this case), does some processing (use the time.Sleep function for the time being),
	and before using the resource, checks if it is required by the other process (the spouse in this case).
 -> If the spouse is hungry, leave the lock on the resource (the spoon) and start demanding the resources again.
	Launch the two goroutines that are passing WaitGroup and the resource’s pointer reference.
 ->For locking and unlocking, use the provided code snippet.

For the output,
try to print something that indicates that spouse one is picking up resource x, checking if the other spouse is hungry, and leaving the resource.
*/

func main() {
	runtime.GOMAXPROCS(4)
	type value struct {
		sync.Mutex
		id     string
		value1 string
		value2 string
		locked bool
	}

	lock := func(v *value) {
		v.Lock()
		v.locked = true
	}

	unlock := func(v *value) {
		v.Unlock()
		v.locked = false
	}

	move := func(wg *sync.WaitGroup, v1, v2 *value) {
		defer wg.Done()
		for i := 0; i <= 3; i++ {
			if i >= 3 {
				fmt.Println("cancelling goroutine...")
				return
			}
			fmt.Printf("%v: %v \n", v1.id, v1.value1)
			lock(v1)
			time.Sleep(2 * time.Second)

			fmt.Printf("%v: Checking if %v \n", v1.id, v2.value2)
			if v2.locked {
				fmt.Printf("%v: Leaving spoon \n", v1.id)
				unlock(v1)
				continue
			}
		}
	}

	a, b := value{
		id:     "Alice",
		value1: "I am picking up spoon",
		value2: "Alice is hungry",
	}, value{
		id:     "Bob",
		value1: "I am picking up spoon",
		value2: "Bob is hungry",
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go move(&wg, &a, &b)
	go move(&wg, &b, &a)
	wg.Wait()
}

/*
Explanation
First, let’s understand the functions of the lock and unlock. The lock and unlock functions are the anonymous functions we use on the locking mechanism.
We use both functions to pass the pointer reference of the resource, which creates and releases the lock. To keep track of this, we use the struct function.

We have one more anonymous function, move, which has an infinite loop that demonstrates livelock.
The first process (Alice) acquires the lock on resource x (the spoon). Before using the resource, Alice checks if the second process (Bob)
requires the resource based on whether Bob is hungry or not. If Bob requires the resource, Alice will release the lock on the spoon.

The important lines in the code are as follows:

Line 54: The first process (Alice) acquires the resource (spoon).
Line 55: The CPU goes to another goroutine for execution.
Line 68: We check if the other process requires the resources (if Bob is hungry).
Line 60: The first process releases the lock on the acquired resource (the spoon).
The reason for livelock in this case is the delay provided by the Sleep function. If we’re executing some heavy computational function, instead of sitting idle, preemption will happen to avoid starvation. The compiler will work on another goroutine and start its execution.

We use WaitGroup to tell the main goroutine to stop until all the goroutines complete their execution.
*/
