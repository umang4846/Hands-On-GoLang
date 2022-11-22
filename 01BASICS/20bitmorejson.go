package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	//encodeJson()
	decodeJson()
}

func encodeJson() {
	lcoCourses := []courses{
		{"ReactJs Bootcamp", 299, "LCO", "123abc", []string{"react", "web"}},
		{"MERN Fullstack", 399, "LCO", "4gfvncff", []string{"full-stack", "web"}},
		{"Angular Bootcamp", 199, "LCO", "bvh344", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "coursename": "ReactJs Bootcamp",
                "Price": 299,
                "website": "LCO",
                "tags": ["react","web"]
        }
   `)
	var lcoCourse courses
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JOSN was valid")
		err := json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is: %v value is: %v and type is: %T\n", k, v, v)
	}
}
