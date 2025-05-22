package main

import (
	"fmt"
)
//function that takes in two integers and returns the sum as an integer
func add(x, y int) int{
	return x+y
}

func main(){
	fmt.Println(add(22,13))
}