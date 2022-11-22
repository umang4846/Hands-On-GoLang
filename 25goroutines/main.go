package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //pointers usually
var signals = []string{"test"}
var mutex = sync.Mutex{}

func main() {
	//go greeter("Hello")
	//greeter("World")
	startTime := time.Now()
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://facebook.com",
		"https://twitter.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
	fmt.Printf("took %v\n", time.Since(startTime))

}

//func greeter(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(3 * time.Millisecond)
//		fmt.Println(s)
//	}
//}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPs in endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
	}
}
