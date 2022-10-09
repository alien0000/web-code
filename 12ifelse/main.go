package main

import "fmt"

func main() {
	fmt.Println("If else in go lange")

	longInCount := 9
	if longInCount < 10 {
		fmt.Println("reguler")
	} else if longInCount > 10 {
		fmt.Println("watch out")
	} else {
		fmt.Println("exetly 10")
	}

	if 7%2 == 0 {
		fmt.Println("even number")
	} else {
		fmt.Println("odd number")
	}

	if num := 3; num < 10 {
		fmt.Println("number is less then 10")
	} else {
		fmt.Println("number is grater then 10")
	}
}
