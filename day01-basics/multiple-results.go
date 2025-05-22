package main

import "fmt"
// function that takes in two string inputs and outputs them as swapped string words
func swap (x,y string)(string, string){
	return y,x
}

func main(){
	a, b := swap("hello" , "world")
	fmt.Println(a,b)
}