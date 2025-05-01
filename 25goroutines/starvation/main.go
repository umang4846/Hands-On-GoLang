package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Overview of starvation
The concepts of concurrency and race condition work fine only when the whole block of code executes without preemption.
But, why is preempting necessary? Can’t the process execute completely and then schedule another process?

Let’s consider a case in which we have five processes. Each process has a different priority, and a preemptive scheduling strategy is being
used to schedule the allocation of the CPU to these processes. The process with the highest priority will execute first.

But what happens if the process with the highest priority keeps executing? The processes with the lower priorities will never get the CPU
to complete their tasks. This situation is known as starvation. A starvation condition occurs when high-priority tasks continue to execute
one after the other, while low-priority tasks starve for a chance to execute.

Preemptive scheduling isn’t the only cause of starvation. It can be caused by some other conditions as well, some of which are as follows:

 -> When a process is never given the resources it needs for execution due to poor resource allocation decisions.
 -> When there aren’t enough resources to support each process as needed.
 -> When a random selection of processes occurs, a process might have to wait for a long time due to non-selection.

In a general, starvation usually means that one or more greedy concurrent processes are unfairly stopping one or more concurrent processes
from completing work efficiently.

Code explanation

In the code above, each function holds of the lock for 10 seconds but in a different sequence.

->greedyRoutine1 holds the lock for 10 seconds in one attempt.

-> greedyRoutine2 holds the lock for 10 seconds as well but not in one attempt. The routine holds for 5 seconds, releases the lock,
	and holds again for 5 seconds.

-> nonGreedyRoutine holds the lock for 10 seconds. The routine holds the lock for 1 second and then releases the lock before asking for the lock again.

From the output of the code widget above, we observe that greedyRoutine1 holds the lock the maximum number of times.
The function greedyRoutine2 also holds the lock about the same number of times but not equal to greedyRoutine1 while the
function nonGreedyRoutine holds the lock the least number of times.

We can observe how two goroutines are greedily holding of the lock and starving the non-greedy routine, though all of the three routines
hold the lock for 10 seconds in total.
*/

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 10 * time.Second

	greedyRoutine1 := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(10 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker 1 was able to execute %v work loops\n", count)
	}

	greedyRoutine2 := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(5 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(5 * time.Nanosecond)
			sharedLock.Unlock()

			count++
		}
		fmt.Printf("Greedy worker 2 was able to execute %v work loops\n", count)
	}

	nonGreedyRoutine := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			for i := 0; i < 10; i++ {
				sharedLock.Lock()
				time.Sleep(1 * time.Nanosecond)
				sharedLock.Unlock()
			}
			count++
		}
		fmt.Printf("Non greedy worker was able to execute %v work loops.\n", count)
	}

	wg.Add(3)
	go greedyRoutine1()
	go greedyRoutine2()
	go nonGreedyRoutine()
	wg.Wait()
}
