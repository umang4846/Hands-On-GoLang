package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//  var myString string
	//  fmt.Scanln(&myString)
	//  fmt.Println(myString)

	// var name string = "Umang"
	// var a, _ = fmt.Println(name)
	// fmt.Println(a)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your full name: ")
	// myName, _ := reader.ReadString('\n')
	// fmt.Println(myName)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your rating: ")
	myrating, _ := reader.ReadString('\n')
	myNumrating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(myNumrating + 2)

}
