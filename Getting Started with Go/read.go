package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type name struct {
		fname string
		lname string
	}
	var arr []name

	var s string
	fmt.Println("Enter the name of text file:")
	fmt.Scanln(&s)
	file, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	for true {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		newName := strings.Split(string(line), " ")
		nameStr := name{
			newName[0], newName[1],
		}
		arr = append(arr, nameStr)
	}
	for _, x := range arr {
		fmt.Println("First Name : " + x.fname)
		fmt.Println("Last Name  : " + x.lname)
	}

}
