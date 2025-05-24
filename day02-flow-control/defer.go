package main

import (
	"fmt"
)
// defer keyword is used to tell a function 
//to hold on execution of a certain part until another section is computed
func main (){
	defer fmt.Println("world")
	fmt.Println("Hello")
}