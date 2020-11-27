package main

import (
	"fmt"
	"strconv"
)

// Swap function to swap two int
func Swap(i int, slice []int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

// BubbleSort func to sort a slice
func BubbleSort(slice []int) {
	// fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				Swap(j, slice)
			}
		}
	}
}

func main() {
	i := 0
	var inp string
	var slice []int
	for true {
		fmt.Println("Enter an integer and Enter X to exit.")
		fmt.Scanln(&inp)
		new, _ := strconv.Atoi(inp)
		if inp == "X" {
			fmt.Println("Exiting....")
			break
		}
		slice = append(slice, new)
		i++
		if i == 10 {
			break
		}
	}
	BubbleSort(slice)
	fmt.Println(slice)
}
