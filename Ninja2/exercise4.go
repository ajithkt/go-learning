package main

import (
	"fmt"
)

func main() {

	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	//Shifting on bit to the left
	y := x << 1
	fmt.Printf("%d\t%b\t%#x", y, y, y)
}
