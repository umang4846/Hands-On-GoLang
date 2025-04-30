package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

/*
Problem :

Write an HTTP service with two endpoints
- First one is going to accept 5 inputs int1, int2, limit, str1, str2 and it will return values
    from 1 to limit, where the values that are multiple of int1 are replaced with str1,
    values that are multiple of int2 are replaced with str2, and
    values that are multiple of int1 and int2 are replaced with str1str2.

- Second endpoint is not going to expect any input, but it will return the input using which
    first endpoint was called the most.

Input Calls to first endpoint:
3, 5, 100, a, d
3, 5, 100, a, d
3, 5, 100, a, d
3, 5, 100, a, d
3, 5, 100, a, d
3, 5, 100, a, b
3, 5, 100, b, c


Output from second endpoint:
3, 5, 100, a, d
5

*/

type Response struct {
	Number int
	Str    string
}

var mu sync.Mutex
var stats = make(map[string]int)

func main() {

	http.HandleFunc("/first", multiplierHandler)
	http.HandleFunc("/stats", statsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func multiplierHandler(w http.ResponseWriter, r *http.Request) {

	int1Str := r.URL.Query().Get("int1")
	int2Str := r.URL.Query().Get("int2")
	limitStr := r.URL.Query().Get("limit")
	str1 := r.URL.Query().Get("str1")
	str2 := r.URL.Query().Get("str2")

	int1, err1 := strconv.Atoi(int1Str)
	int2, err2 := strconv.Atoi(int2Str)
	limit, err3 := strconv.Atoi(limitStr)

	if err1 != nil || err2 != nil || err3 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var res []Response
	for i := 1; i <= limit; i++ {
		if i%int1 == 0 && i%int2 == 0 {
			res = append(res, Response{Number: i, Str: str1 + str2})
		} else if i%int1 == 0 {
			res = append(res, Response{Number: i, Str: str1})
		} else if i%int2 == 0 {
			res = append(res, Response{Number: i, Str: str2})
		}
	}

	key := fmt.Sprintf("%d,%d,%d,%s,%s", int1, int2, limit, str1, str2)
	mu.Lock()
	stats[key]++
	mu.Unlock()

	json.NewEncoder(w).Encode(res)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	maxCount := 0
	maxFreqKey := ""

	for key, value := range stats {
		if value > maxCount {
			maxFreqKey = key
			maxCount = value
		}
	}

	res := map[string]interface{}{
		"Key":   maxFreqKey,
		"Value": maxCount,
	}

	json.NewEncoder(w).Encode(res)
}
