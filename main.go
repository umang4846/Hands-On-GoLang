package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	basketID = "6ab24992434d5d2ba9700bf76d387208"
)

func main() {
	url := "https://pp-stress.pragmaticplaysports.net/betbasket/"
	method := "POST"

	payload := []byte(`{"Id":"","currency":"USD"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set headers
	req.Header.Add("accept", "application/json")
	req.Header.Add("bb-expires", "Mon, 12 Aug 2024 07:09:40 GMT")
	req.Header.Add("bb-session-id", "405647f6beabda6f51e3ad63650b7ef0")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

	fmt.Println("Starting next request")
	makeSecondRequest()
}

func makeSecondRequest() {
	url := "https://pp-stress.pragmaticplaysports.net/betbasket/6ab24992434d5d2ba9700bf76d387208/single"
	method := "POST"

	payload := []byte(`{"marketId":"BG-36712974","price":{"up":1,"down":1,"dec":"2"},"side":"AWAY","line":0,"selectionids":["W2-2"]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("ERROR AFTER NEW REQUEST", err)
		return
	}

	// Set headers
	req.Header.Add("accept", "application/json")
	req.Header.Add("bb-expires", "Mon, 12 Aug 2024 07:15:25 GMT")
	req.Header.Add("bb-session-id", "896ca7c78dacf00c7bed4cbe73c8209b")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("ERROR AFTER CLIENT DO", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
