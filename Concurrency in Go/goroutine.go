package main

import "fmt"

/*
Raca condition is when multiple threads are trying to access and manipulat the same variable.
In the code below, main, addone and print are all accessing and changing the value of x.
Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable.
*/
var pt int = 1

func print() {
	fmt.Println(pt)
}

func addone() {
	pt = pt + 1
	fmt.Println(pt)
}
func main() {
	i := 0
	go addone()
	go print()

	i++
	fmt.Println()

}
