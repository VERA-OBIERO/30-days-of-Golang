package main

import (
	"fmt"
	"math"
)
// this is a function that takes in a float integer x
// checks if it is a negative number and squares it then adds a complex imaginary number and returns it
// if it is positive, it finds the square root and returns it a as a string
func sqrt( x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))// fmt.Sprint converts output to string
}

func main() {
	fmt. Println(sqrt(2), sqrt(-4))
}