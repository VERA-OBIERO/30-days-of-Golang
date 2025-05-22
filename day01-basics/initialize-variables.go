package main

import "fmt"
//when variables are initialized with a value then type is not mandatory
//the variable automatically takes the type of the value assigned to it
var i , j = 1, 2

func main(){
	var c, java, python = 1, true, "Hi"
	fmt.Println(i, j, c, java, python)
}