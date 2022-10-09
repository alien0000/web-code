package main

import "fmt"

func main() {
	defer fmt.Println("hii")
	fmt.Println("welcome to defer in go lang")
	fmt.Println("world")
	num()
}

func num() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
