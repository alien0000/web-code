package main

import "fmt"

func main() {
	fmt.Println("welcome to function in go lang")
	greed()
	restul := addTowNum(3, 3)
	fmt.Println("the result is: ", restul)
	fmt.Println(proAdder(1, 2, 3, 4))
}
func greed() {
	fmt.Println("hii.. there from go developers")

}

func addTowNum(numOne int, numTow int) int {
	return numOne + numTow
}
func proAdder(values ...int) (int, string) {
	total := 0
	for _, k := range values {
		total += k
	}
	return total, "this is message from proadder"
}
