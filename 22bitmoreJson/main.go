package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platfrom string   `json:"wensite"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome to json video")
	// EncodeJson()
	DecodeJson()
}
func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs", 200, "google.com", "asdfg", []string{"web", "js"}},
		{"Js", 200, "js.com", "012369", []string{"web", "javascript"}},
		{"ReactJs", 200, "google.com", "7885dfg", nil},
	}

	// package this data as JSON data
	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	// finalJson, err := json.MarshalIndent(lcoCourses, "sajda", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs",
		"Price": 200,
		"wensite": "google.com",
		"tags": ["web","js"]
}
	`)
	var lcoCourse course

	checkValid := json.Valid((jsonDataFromWeb))
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was NOT valid")
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is :%T\n", k, v, v)
	}
}
