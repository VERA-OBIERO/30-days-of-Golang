package main

import(
	"fmt"
	"math"
)
// this is a function that takes in three float integers
// it returns a float
// the first two integers are computed as a power and assigned to a new variable
// then the new variable is compared to the third integer
// if condition is true, new varaible is reurned, else third integer is returned.
func pow(x, n, val float64) float64{
	if v := math.Pow(x, n ) ; v < val {
		return v
	}
	return val
}

func main () {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}