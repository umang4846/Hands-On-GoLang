package main

import "fmt"

type User1 struct {
	Name  string
	Email string
	Age   int
}

func (u User1) getEmail() {
	fmt.Println("Email is : ", u.Email)
}

func (u User1) NewEmail() {
	u.Email = "test@gmil.com"
	fmt.Println("New Email is : ", u.Email)
}

func main() {
	rob := User1{"rob", "rob@gmail.com", 34}
	fmt.Printf("%v\n", rob)
	fmt.Printf("%+v\n", rob)
	fmt.Println("Name=", rob.Name)
	rob.getEmail()
	rob.NewEmail()
	fmt.Printf("%+v\n", rob)

	var sam = new(User1)
	sam.Name = "sam"
	sam.Email = "sam@gmail.com"
	sam.Age = 29
	fmt.Printf("%+v\n", sam) //will get '&' as initialized with new keyword

	var tobby = &User1{"tobby", "tobby@gmail.com", 34}
	fmt.Printf("%+v", tobby)

}
