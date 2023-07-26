package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"stock-api-go-postgres/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8000...")

	log.Fatal(http.ListenAndServe(":8000", r))
}
