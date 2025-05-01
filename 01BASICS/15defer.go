package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()

	err := readFileHere("output.txt")
	if err != nil {
		fmt.Println(err)
	}
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func readFileHere(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil

}
