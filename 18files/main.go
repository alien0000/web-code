package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to file handeling in go")
	content := "this is the text its going to the file"
	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)
	fmt.Println("the leng: ", length)
	defer file.Close()
	reader("./myfile.txt")
}

func reader(file string) {
	databy, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("the line is :\n", string(databy))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
