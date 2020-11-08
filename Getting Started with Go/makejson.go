package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	map1 := make(map[string]string)
	var s string
	fmt.Println("Enter your name:")
	fmt.Scanln(&s)
	map1["name"] = s
	fmt.Println("Enter your address:")
	fmt.Scanln(&s)
	map1["address"] = s
	fmt.Println(map1)

	obj, err := json.Marshal(map1)
	if err == nil {
		fmt.Println(string(obj))
	}
}
