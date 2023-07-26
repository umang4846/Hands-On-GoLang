package main

import (
	"encoding/json"
	"fmt"
)

/*func main() {
	fmt.Println("Hello, Reader! Your learning about 'goto' statement")
	// We create a for loop which runs until i is 10
	for i := 0; i < 10; i++ {
		fmt.Printf("Index: %d\n", i)
		if i == 5 {
			// When i is 5, lets exit by using goto
			goto exit
		}
	}
	fmt.Println("Skip this line here")
	// Create the exit label and insert code that should be executed when triggered
exit:
	fmt.Println("We are now exiting the program")
}*/
/*func main() {
	name := "Alice"
	age := 30
	info := fmt.Sprintf("My name is %s and I am %d years old.\n", name, age) // formatted string creation

	var readName string
	var readAge int
	_, _ = fmt.Sscanf(info, "My name is %s and I am %d years old.\n", &readName, &readAge) // read from the formatted string

	fmt.Println(readName, readAge)

	alice := Person{"Alice", 30}
	fmt.Printf("%v\n", alice)  // {Alice 30}
	fmt.Printf("%+v\n", alice) // {Name:Alice Age:30}
	fmt.Printf("%#v\n", alice) // main.Person{Name:"Alice", Age:30}
	fmt.Printf("%T\n", alice)  // main.Person

	fmt.Printf("%t\n", true) // true
	fmt.Printf("%b\n", 5)    // 101
	fmt.Printf("%c\n", 65)   // A

	fmt.Printf("%e\n", 123400000.0) // 1.234000e+08
	fmt.Printf("%f\n", 123.456789)  // 123.456789
}*/

func main() {
	// Start tracing
	/*traceFile, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer traceFile.Close()

	err = trace.Start(traceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	// Your code to trace goes here...

	// For example, simulate some work
	for i := 0; i < 10; i++ {
		_ = doWork(i)
	}*/

	bb := "{\n    \"id\": \"ebaaea18effd983cb5d8881345104749\",\n    \"playerId\": \"assa\",\n    \"brandId\": \"pp\",\n    \"singles\": [\n        {\n            \"singleId\": \"U-8349582.W2-1\",\n            \"legType\": \"SIMPLE\",\n            \"sportId\": \"BOX\",\n            \"sportName\": \"\",\n            \"competitionId\": \"U-1366\",\n            \"competitionName\": \"\",\n            \"eventId\": \"U-147818\",\n            \"eventName\": \"Crews, Franchon v Marshall, Savannah\",\n            \"eventStatus\": \"OFFERED\",\n            \"eventTradeStatus\": \"TRADABLE\",\n            \"marketId\": \"U-8349582\",\n            \"marketName\": \"Match Betting\",\n            \"marketStatus\": \"OFFERED\",\n            \"marketTradeStatus\": \"TRADABLE\",\n            \"stake\": \"0\",\n            \"potReturn\": \"0\",\n            \"promoReturn\": \"0\",\n            \"bunker\": false,\n            \"price\": {\n                \"up\": 7,\n                \"down\": 2,\n                \"dec\": \"4.5\"\n            },\n            \"side\": \"HOME\",\n            \"line\": 0,\n            \"eachWay\": {\n                \"placeReduction\": \"0\"\n            },\n            \"placeReduction\": \"0\",\n            \"interrelated\": null,\n            \"selections\": [\n                {\n                    \"id\": \"W2-1\",\n                    \"participantId\": \"U-325939\",\n                    \"name\": \"Crews, Franchon\",\n                    \"status\": \"OFFERED\",\n                    \"tradeStatus\": \"TRADABLE\",\n                    \"price\": {\n                        \"up\": 7,\n                        \"down\": 2,\n                        \"dec\": \"4.5\"\n                    },\n                    \"side\": \"HOME\",\n                    \"line\": 0,\n                    \"eachWay\": {\n                        \"isEachWay\": false,\n                        \"placeReduction\": \"0\"\n                    },\n                    \"pushHonored\": false,\n                    \"homeScore\": 0,\n                    \"awayScore\": 0\n                }\n            ],\n            \"status\": \"INCLUDED\",\n            \"confirmedStatus\": \"INCLUDED\",\n            \"available\": true,\n            \"cappingLimit\": \"0\",\n            \"multipleDisabled\": false,\n            \"isFreeBet\": false,\n            \"freeBetAmount\": \"0\",\n            \"bonusId\": 0,\n            \"sbBonusAwardId\": 0\n        }\n    ],\n    \"currency\": \"EUR\",\n    \"stake\": \"0\",\n    \"potReturn\": \"0\",\n    \"updatedAt\": \"2023-06-27T08:39:59.072691621Z\",\n    \"status\": \"OPEN\"\n}"

	b := []byte(bb)

	var objmap map[string]interface{}

	if err := json.Unmarshal(b, &objmap); err != nil {
		panic(err)
	}
	fmt.Println(objmap)

	singles := objmap["singles"]
	fmt.Printf("singles: %T", singles)

	switch t := singles.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case []interface{}:

		var keys []string

		if key, ok := singles.(string); ok {
			keys = append(keys, key)
		}

		fmt.Println(keys)
	}

	v, ok := singles.(map[string]interface{})

	if !ok {
		fmt.Println("error ", ok)
	}
	fmt.Println("v:", v)
	price := v["price"]
	fmt.Println("price:", price)

	//single := earnings["price"].(interface{})

}

func doWork(i int) int {
	// Simulate some work
	result := i * 2

	return result
}
