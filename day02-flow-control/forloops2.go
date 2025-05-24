package main

import "fmt"
// this is a for loop that initalizes sum as 1
// if the value of sum is less than 1000, sum is doubled
// the value of sum is printed when the condition evaluates to false
func main(){
	sum := 1
	for ; sum < 1000 ;{ // we can also write for loops by omiitting the init and post statements
		sum += sum
	}
	fmt.Println(sum)
}