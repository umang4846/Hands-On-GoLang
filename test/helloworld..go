package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorld2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World Started")

	//sleep for some second

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Hello World Ended")

}
