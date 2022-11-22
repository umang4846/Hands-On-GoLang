package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	rob := User{"rob", "rob@gmail.com", 34}
	fmt.Printf("%v\n", rob)
	fmt.Printf("%+v\n", rob)
	fmt.Println("Email is ", rob.Email)

	var sam = new(User)
	sam.Name = "sam"
	sam.Email = "sam@gmail.com"
	sam.Age = 29
	fmt.Printf("%+v\n", sam) //will get '&' as initialized with new keyword

	var tobby = &User{"tobby", "tobby@gmail.com", 34}
	fmt.Printf("%+v", tobby)
}
