package main

import "fmt"
// variable declarations can happen outside or inside a function and must be assigned a type
var c, python, java bool

func main(){
	var i int
	i=2
	fmt.Println(i, c, python, java)// varibles are automatically 0 or false if no value is assigned to them
}