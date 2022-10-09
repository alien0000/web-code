package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")
	sajda := User{"sajda", "sajda@go.dev", true, 16}
	fmt.Println(sajda)

	fmt.Printf("sajda details are: %v\n", sajda)
	fmt.Printf("sajda details are: %+v\n", sajda)
	fmt.Printf("user name are: %v\n", sajda.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
