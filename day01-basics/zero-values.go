package main

import "fmt"
// variables declared without any value are automatically assigned 0, false or empty string
// constants van only be declared using const keyword
func main() {
	var i int
	var f float64
	var b bool
	var s string
	const y = 7
	fmt.Printf("%v %v %v %q %v\n", i, f, b, s, y)
}