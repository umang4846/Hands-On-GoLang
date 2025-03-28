package main

import "fmt"

func main() {
	str := "Hello"

	newString := ""

	for i := 0; i < len(str); i++ {
		newString = string(str[i]) + newString
	}
	fmt.Println("reversed string :" + newString)
}
