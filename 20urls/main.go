package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=cmkwnf65"

func main() {
	fmt.Println("welcom to url in go lang")
	fmt.Println(myurl)
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("the type is : %T\n ", qparams)

	for _, vale := range qparams {
		fmt.Println(vale)
	}

	partsOfUrl := &url.URL{
		Scheme:  "htpps",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=sajda",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
