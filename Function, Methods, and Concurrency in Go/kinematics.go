package main

import (
	"fmt"
	"math"
)

// GenDisplacementFn gives displacement wrt time
func GenDisplacementFn(a, v, s float64) func(float64) float64 {

	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v*t + s
	}
}
func main() {
	var a, v, s float64

	fmt.Println("Enter acceleration")
	fmt.Scanln(&a)
	fmt.Println("Enter velocity")
	fmt.Scanln(&v)
	fmt.Println("Enter initial displacement")
	fmt.Scanln(&s)

	fn := GenDisplacementFn(a, v, s)
	var t float64
	fmt.Println("Enter time")
	fmt.Scanln(&t)
	fmt.Println(fn(t))

}
