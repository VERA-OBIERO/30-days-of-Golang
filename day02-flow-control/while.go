package main 

import "fmt"
//in Go, while loops are written using the for keyword
// this function initializes sum as 1 and while it is less than 1000, sum is doubled.
// in this syntax, the semicolons for init and post statements are ommited.
func main(){
	sum := 1
	for sum < 1000{
		sum += sum
	}
	fmt.Println(sum)
}