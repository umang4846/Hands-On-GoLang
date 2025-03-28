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

	//assign function as variable and then Call it using the variable
	x := print // assign function to the variable
	x(5)       // call the function using that varible and we can pass the arguments as well.

	test := func(x int) {
		fmt.Println("Hello! function is assigned to variable ", x)
	}
	test(5)

	test1 := func(x int) int {
		fmt.Println("Hello! function is returning something ", x)
		return x * -1
	}(8)
	fmt.Println(test1)

	// pass function to another function
	test2 := func(x int) int {
		fmt.Println("Function is passed to another function"+
			" ", x)
		return x * -1
	}
	testFunc(test2)

	//function closure
	returnFunc("hello")()

	closer := returnFunc("goodbye")
	closer()
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

func print(x int) {
	fmt.Println("Hello ", x)
}

func testFunc(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}
