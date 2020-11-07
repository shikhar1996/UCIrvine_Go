package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var arr []int = make([]int, 3)
	var inputString string

	for true {
		fmt.Println("Enter an integer to continue.\nEnter 'exit' to close.")
		fmt.Scanln(&inputString)

		if inputString == "exit" {
			break
		}
		number, err := strconv.Atoi(inputString)
		if err != nil {
			fmt.Println("Input Error")
			continue
		}
		arr = append(arr, number)

		sort.Ints(arr[3:])
		fmt.Println(arr[3:])
	}

}
