package main

import (
	"fmt"
)

func main() {
	x := [5]int{}
	fmt.Println(x)
	x[0] = 5
	x[4] = 1
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)

}
