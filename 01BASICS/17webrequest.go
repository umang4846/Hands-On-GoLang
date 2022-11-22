package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const myUrl = "https://lco.dev"

func main() {
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()
	datbytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(datbytes)
	fmt.Println(content)
}
