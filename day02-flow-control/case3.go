package main 

import (
	"fmt"
	"time"
)
// this is a switch function that checks the time and prints the appropriate greeting
// it is not mandatory for switch statements to have a condition
func main (){
	clock := time.Now()
	switch  {
	case clock.Hour() < 12:
		fmt.Println("Good morning!")
	case clock.Hour() < 17 :
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")

	}
}