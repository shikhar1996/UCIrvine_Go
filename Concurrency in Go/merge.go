package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func subSort(wg *sync.WaitGroup, subArr []int) {
	sort.Ints(subArr)
	// fmt.Println(subArr)
	wg.Done()
}

func main() {
	fmt.Println("Hello, please enter integers to sort seperated by blanks.")
	br := bufio.NewReader(os.Stdin)
	temp, _, _ := br.ReadLine()
	numbers := strings.Split(string(temp), " ")
	fmt.Println(numbers)
	var arr []int
	for _, s := range numbers {
		n, _ := strconv.Atoi(s)
		arr = append(arr, n)
	}
	fmt.Println(arr)

	size := len(arr) / 4
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		if i != 3 {
			go subSort(&wg, arr[i*size:(i+1)*size])
		} else {
			go subSort(&wg, arr[i*size:])
		}
	}
	wg.Wait()
	s1 := arr[:size]
	s2 := arr[size : 2*size]
	mergedList := append(s1, s2...)
	s3 := arr[2*size : 3*size]
	mergedList = append(mergedList, s3...)
	s4 := arr[3*size:]
	mergedList = append(mergedList, s4...)
	fmt.Println("Sorted list of Integers is:")
	wg.Add(1)
	go subSort(&wg, mergedList)
	wg.Wait()
	fmt.Println(mergedList)
}
