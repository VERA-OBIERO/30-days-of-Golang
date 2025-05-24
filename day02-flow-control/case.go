package main

import (
	"fmt"
	"runtime"
)
// a swith function that checks what opearting system a machine is using
func main () {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		//freebsd, openbsd
		// plan9, windows...
		fmt.Printf("%s. \n", os)
	}
}