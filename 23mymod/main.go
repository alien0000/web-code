package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello mod in go lang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("hey there mo user")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to golang series on youtube</h1>"))
}

// go mod tidy
// go mod verify --> to verify the requirements
// go mod list
// go mod list all
// go list -m all
//  go list -m -versions github.com/gorilla/mux  --> to list all the version

// go mod why github.com/gorilla/mux

// go mod graph
// go mod edit -go 1.17
// go mod edit -module 1.16
// go mod vendor  --> if the modules not working run this it's gies a vesdor folder

// go run -mod=vendor main.go  --> it's 1st lookup in the vendor folder
