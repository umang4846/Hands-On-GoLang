package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Handing simple text files in Go!")
	content := "this text will be saved in TXT file"
	file, err := os.Create("./creatfile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./creatfile.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Content \n", string(databyte))
}
