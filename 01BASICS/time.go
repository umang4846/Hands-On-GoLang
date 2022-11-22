package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning time")
	presentTime := time.Now()
	fmt.Println("Present time is : ", presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.April, 27, 25, 61, 55, 0, time.UTC)
	fmt.Println(createdDate)

}
