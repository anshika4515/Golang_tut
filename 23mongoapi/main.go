package main

import (
	"fmt"
	"log"
	"main/router"
	"net/http"
)

func main() {
	fmt.Println("Learn abt mongo")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listing at port 4000")
}
