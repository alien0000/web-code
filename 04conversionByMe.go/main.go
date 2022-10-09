package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hi there ... give give us ratting for our pizza")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reader)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
