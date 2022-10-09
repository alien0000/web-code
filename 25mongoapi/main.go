package main

import (
	"fmt"
	"github/sajdakabir/mongoapi/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("mongoDB in go lang")
	r := router.Router()
	fmt.Println("server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listing at port 4000...")

}
