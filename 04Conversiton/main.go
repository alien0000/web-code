package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to oure pizza app")
	fmt.Println("plase rete out pizza")
	reader := bufio.NewReader((os.Stdin))
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to youre rattiong", numRating+1)
	}
}
