package main

import (
	"fmt"
	//"math"
)
// this is a function that computes the squareroot of a given number x without using Sqrt import from math
// we instead use a formula
func squareroot (x float64) float64 {
	z := 1.0
	for i := 0; i <10 ; i++ {
		fmt.Println("Step", i, "z =", z)
		z = z - ((z*z - x) / (2*z)) // Newton's method for computing square root
	}
	return z
}

func main(){
	fmt.Println(squareroot(2))
}