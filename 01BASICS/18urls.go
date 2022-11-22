package main

import (
	"fmt"
	"net/url"
)

const lcoUrl = "https://lco.dev:3000/learn?course=reactjs&paymentid=35jnrfcnjndv"

func main() {
	fmt.Println(lcoUrl)
	result, _ := url.Parse(lcoUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params are : %T\n", qparams)

	fmt.Println(qparams["course"])

	for _, v := range qparams {
		fmt.Println("Param is : ", v)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh"}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
