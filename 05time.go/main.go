package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006  15:04:03 Monday"))

	createData := time.Date(2020, time.January, 01, 23, 24, 0, 0, time.UTC)
	fmt.Println(createData)
	fmt.Println(createData.Format("01-02-2006 Monday"))
}
