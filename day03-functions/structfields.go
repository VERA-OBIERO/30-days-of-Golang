package main

import "fmt"
// how to access and change elemnents in a structure
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 5
	fmt.Println(v)
}