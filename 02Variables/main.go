package main

import "fmt"

// jwtToken := 300000   gobally not allowed

const LoginToken string = "bsiuabfwjdbcdfsk" //public

func main() {
	var username string = "sajda"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	// find "the go programming lan specification " -->>type section
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45141848995655548798
	// var smallFloat float32 = 255.45141848995655548798
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("variable is of type: %T \n", anotherString)

	// implicit type

	var website = "google.com"
	fmt.Println(website)

	// no var style

	numberOfUser := 30000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}
