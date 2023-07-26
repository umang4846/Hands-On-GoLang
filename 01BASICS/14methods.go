package main

import (
	"fmt"
	"golang.org/x/tools/go/analysis/passes/fieldalignment"
	"golang.org/x/tools/go/analysis/singlechecker"
	"net/url"
	"strings"
	"unsafe"
)

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

func (u *User1) SetEmail(email string) {
	fmt.Printf("into method %p\n", u)
	u.Email = email
}

func main() {
	/*rob := User1{"rob", "rob@gmail.com", 34}
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

	u1 := User1{
		Name:  "Name",
		Email: "Email",
		Age:   1,
	}
	u := &u1
	fmt.Println(u)
	fmt.Printf("%p\n", &u)
	u.SetEmail("emailNew")
	fmt.Println(u)
	fmt.Printf("%p\n", &u)*/
	singlechecker.Main(fieldalignment.Analyzer)
	type testStruct struct {
		testFloat1 float64
		testFloat2 float64
		testBool1  bool
		testBool2  bool
	}

	a := testStruct{}
	fmt.Println(unsafe.Sizeof(a)) // 24 bytes

	url1 := "/betbasket/539616e8a60d2e0fc03e52b81babb28d/multimarket/betbuilder-BG-11722180.W2-1&&WS2-12-Y%7CBG-11722121.W2-1%7CBG-11722140.H2-1%7CBG-11721348.W2-1D%7CBG-11722232.W2-1%7CBG-11722178.W2-1D"
	//url2 := "https://dev-ppbet-dazn-es.uatsecure.com/betbasket/78ee9b549f62837c457827f77b83f2f9/multimarket/betbuilder-BG-11721752.W2-1|BG-11722087.W2-1"

	decoded1, err := url.PathUnescape(url1)
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return
	}

	decoded2, err := url.PathUnescape(url1)
	if err != nil {
		// Handle error
		fmt.Println("Error decoding URL:", err)
		return
	}

	if decoded1 == decoded2 {
		fmt.Println("Its same****")
	}
	fmt.Println("Decoded value:", decoded1)
	fmt.Println("Decoded value:", decoded2)

	fmt.Println("NEW URL:", normalizePath(url1))
}

func normalizePath(path string) string {
	// Split the path into segments
	segments := strings.Split(path, "/")

	// Remove empty segments and collapse consecutive slashes
	cleanSegments := make([]string, 0, len(segments))
	for _, segment := range segments {
		if segment != "" {
			cleanSegments = append(cleanSegments, segment)
		}
	}

	// Reconstruct the normalized path with a single leading slash
	normalizedPath := "/" + strings.Join(cleanSegments, "/")

	return normalizedPath
}
