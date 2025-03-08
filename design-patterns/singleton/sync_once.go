package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
	//empty struct
}

var singleInstance1 *single

func getInstance1() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &Single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance1
}
