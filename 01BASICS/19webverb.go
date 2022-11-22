package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web server")

	//performGetRequest()
	//performPostRequest()
	performPostFormRequest()

}

func performPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "umang")
	data.Add("lastname", "patel")
	data.Add("email", "umang@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func performPostRequest() {
	const myurl = "http://localhost:8000/post"
	requestBody := strings.NewReader(
		`{
		"coursename":"Let's go with GO",
		"price":"0",
		"platform":"learncodeonline.in"
	}`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code is :", response.StatusCode)
	fmt.Println("Content length is :", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	bytecount, _ := responseString.Write(content)

	fmt.Println("Byte count is : ", bytecount)
	fmt.Println(responseString.String())

	fmt.Println(content)
	fmt.Println(string(content))
}
