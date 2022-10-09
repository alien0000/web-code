package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to switch and case in go lang")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println("the dice number is ", diceNum)
	switch diceNum {
	case 1:
		fmt.Println("youre number is 1 you can move ")
	case 2:
		fmt.Println("you can move 2 sport")

	case 3:
		fmt.Println("you can move 3 sport")
		fallthrough // it alow to run next line
	case 4:
		fmt.Println("you can move 4 sport")
	case 5:
		fmt.Println("you can move 5 sport")
	case 6:
		fmt.Println("you can move 6 sport , you can chance roll up the dice")
	default:
		fmt.Println("What is this!!")
	}

}
