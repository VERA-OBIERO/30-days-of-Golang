package main

import (
	"fmt"
	"math/cmplx"
)
// Go data types include:
// 1. Boolean i.e true or false
// 2. Strings i.e "bread", "juice"
// 3. Integers could be signed(can be negative) or unsigned(only positive) of bits 8, 16, 32, 64
//     uint8 = byte(raw data byte), int32 = rune (single character in a unicode)
// 4. Floats i.e 3.123444
// 5 Complex number - combination of real and imagined numbers e.g 7+6i
// N.B. Always use int for integer values unless otherwise.
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}