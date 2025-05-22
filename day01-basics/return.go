package main // tells the compiler that this is an executable go program

import "fmt" // formatted output challenge
//a function that takes in an integer, stores it in a variable called sum 
//then splits it into two according to the criteria in the function
func split(sum int)(x, y int){
	x = sum * 4 / 9 // since this is an integer, the decimals are dropped and only whole numbers considered
	y = sum - x
	return // naked return - allows more than one output to be returned
}

func main(){
	fmt.Println(split(17))
}