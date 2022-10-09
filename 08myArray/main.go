package main

import "fmt"

func main() {
	fmt.Println("welcome to class of array")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "mushroom"}
	fmt.Println("vegy list is: ", vegList)
}
