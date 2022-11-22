package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	//front end
	//take user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name")
	name, _ = reader.ReadString('\n')

	//take rating from user and convert it in int
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter rate between 1 to 5")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Backend
	//fmt.Printf("%v, %v", name, mynumRating)
	fmt.Printf("Hello %v, \nThanks for rating our service with %v star.\n\nYour rating was recorded in our system at %v \n", name, mynumRating, time.Now().Format(time.Stamp))

	if mynumRating == 5 {
		fmt.Println("Bonus for team for 5 star")
	} else if mynumRating == 4 || mynumRating == 3 {
		fmt.Println("We are always improving")
	} else {
		fmt.Println("Need Serious improvements")
	}
}
