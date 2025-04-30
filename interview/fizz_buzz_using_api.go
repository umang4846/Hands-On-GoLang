package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	http.HandleFunc("/fizzbuzz", fizzBuzzHandler)
	fmt.Println("Sever is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {

	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	from, err1 := strconv.Atoi(fromStr)
	to, err2 := strconv.Atoi(toStr)

	if err1 != nil || err2 != nil || from > to {
		http.Error(w, "Invalid from or to parameters", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make([]string, to-from+1) // Pre-allocate slice for ordered output

	fmt.Printf("FizzBuzz results for range %d to %d:\n", from, to)

	for i := from; i <= to; i++ {
		wg.Add(1)
		go func(n, index int) {
			defer wg.Done()
			mu.Lock()
			results[index] = fmt.Sprintf("%d: %s", n, fizzBuzzLogic(n))
			mu.Unlock()
		}(i, i-from)
	}
	for _, res := range results {
		fmt.Println(res)
	}

	wg.Wait()
	fmt.Println("All numbers processed.")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Fizzbuzz results printed to console."))

}

func fizzBuzzLogic(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}
