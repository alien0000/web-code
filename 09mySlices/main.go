package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to class on slices")

	var fruiteList = []string{"Apple", "Tomato", "peach"}
	fmt.Printf("Type of fiuiteList is %T\n", fruiteList)
	fmt.Println(fruiteList)
	fruiteList = append(fruiteList, "banana", "mango")
	fmt.Println(fruiteList)
	fruiteList = append(fruiteList[1:])
	fmt.Println(fruiteList)
	fruiteList = append(fruiteList[1:3])
	fmt.Println(fruiteList)

	fruiteList = append(fruiteList[:3])
	fmt.Println(fruiteList)

	highScore := make([]int, 4)
	highScore[0] = 789
	highScore[1] = 59
	highScore[2] = 68
	highScore[3] = 984
	// highScore[4] = 5555
	highScore = append(highScore, 555, 999)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted((highScore)))
	sort.Ints(highScore)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted((highScore)))

	// how to remove a value from slices based on index

	var lang = []string{"c", "c++", "python", "java", "ruby", "go", "javascript"}
	fmt.Println(lang)
	var index int = 3
	lang = append(lang[:index], lang[index+1:]...)
	fmt.Println(lang)
}
