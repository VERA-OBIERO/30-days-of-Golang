package main

import "fmt"
// for loops are separated by semicolons and have 3 constructs
func main(){
	// this is a for loop that initializes sum as 0 and i as 0
	// as long as i is less than 10, i is incremented by 1 and values of i summed then printed.
	sum := 0;
	for i :=0 ; i < 10 ; i++{ // 1. the init statement(executed before 1st iteration);
		                      // 2. the condition expression (evaluated before every iteration) ; 
		sum += i              // 3. the post statement(executed at the end of every iteration);
	}
	fmt.Println(sum)
}
