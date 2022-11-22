package main

import (
	"fmt"
)

func main() {
	superman()
	result := multiply(5, 6)
	fmt.Println(result)
	result = adder(3, 5, 67, 4, 2, 7, 5, 4)
	fmt.Println(result)
	sum, length, name := adder1(3, 4, 5, 6, 7, 88, 8, 8)
	fmt.Println(sum, length, name)
}

func superman() {
	fmt.Println("I am superman")
}

func multiply(v1 int, v2 int) int {
	return v1 * v2
}

func adder(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func adder1(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	length := len(values)
	name := "just for fun"
	return sum, length, name
}
