package main

import "fmt"

func main() {
	fmt.Println("Wlcome to class on Maps")

	lang := make(map[string]string)
	lang["js"] = "javascript"
	lang["py"] = "python"
	lang["rb"] = "ruby"
	fmt.Println(lang)
	fmt.Println(lang["js"])
	delete(lang, "rb")
	fmt.Println(lang)

	for key, value := range lang {
		// fmt.Printf("the value %v, and the key is %v\n ", key, value)
		fmt.Println("key is", key, "value is ", value)
	}
}
