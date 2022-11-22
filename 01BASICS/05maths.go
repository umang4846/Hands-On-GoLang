package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to Maths in golang")
	var myNumberOne int = 2
	var myNumberTwo float64 = 4

	fmt.Println("Sum is : ", myNumberOne+int(myNumberTwo))
	//random number
	//	rand.Seed(time.Now().UnixNano())
	//	fmt.Println(rand.Intn(5) + 1) //random from 1 to 5

	//random from crypto
	myRandomNum, err := rand.Int(rand.Reader, big.NewInt(5))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myRandomNum)
}
