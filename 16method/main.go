package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")
	sajda := User{"sajda", "sajda@go.dev", true, 16}
	fmt.Println(sajda)

	fmt.Printf("sajda details are: %v\n", sajda)
	fmt.Printf("sajda details are: %+v\n", sajda)
	fmt.Printf("user name are: %v\n", sajda.Name)
	fmt.Printf("user mail are: %v\n", sajda.Email)
	sajda.getStatus()
	sajda.getMail()
	fmt.Printf("user mail are: %v\n", sajda.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("the mail is ", u.Status)
}

func (u User) getMail() {
	u.Email = "test@dev.com"
	fmt.Println("the mail is ", u.Email)
}
