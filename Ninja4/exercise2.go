package main

import (
	"fmt"
)

func main() {
	//using slice
	x := []int{10, 12, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)

}
