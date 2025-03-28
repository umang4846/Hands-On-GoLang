package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	//rate allowed per second
	//	NewRateLimiter(10, 1)

	//have the http router

	http.HandleFunc("/", helloWorld2)
	http.ListenAndServe(":8080", nil)
}

/*func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
*/

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World Started")

	//sleep for some second

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Hello World Ended")

}
func helloWorld1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World Started")

	//sleep for some second

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Hello World Ended")

}
