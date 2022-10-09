package main

import "fmt"

func main() {
	fmt.Println("welcome to class of pointer")
	myNumber := 25
	var ptr = &myNumber
	fmt.Println("print the value of ptr ", ptr)
	fmt.Println("valu of actual ptr", *ptr)
	*ptr = *ptr + 2
	fmt.Println("valu of actual ptr", *ptr)
}
