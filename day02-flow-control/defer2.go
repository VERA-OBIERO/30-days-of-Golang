package main 

import (
	"fmt"
)
// defer keyword allows for a LIFO approach which will be helpful when dealing with stacks
func main() {
	fmt.Println("counting")
	for i :=0 ; i < 10 ; i++{
		defer fmt.Println(i)
	}
	fmt.Println("done")
}