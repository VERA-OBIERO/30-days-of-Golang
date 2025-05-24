package main

import (
	"fmt"
	"math"
)
// this is a function that takes in 3 float integers
// it gets the power of the first 2 and assigns them to a new variable y
// then compares y to the 3rd integer
// if true, y is returned... if not a statemnt that val of y > than value of 3rd integer is printed out
func pow(x,n, val float64) float64{
	if y := math.Pow(x,n) ; y < val {
		return y
	}else {
		fmt.Printf("%g >= %g\n", y, val)
	}
	//the variable y that was introduced only exists within the scope of the if else block
	return val
}

func main(){
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}