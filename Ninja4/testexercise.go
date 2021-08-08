package main

import (
	"fmt"
)

func main() {
	// Use slice instead of array
	x := []int{5, 6, 7, 8}
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}
	for j := 0; j < len(x); j++ {
		fmt.Printf("%v,%v\n", j, x[j])
	}
	//Slicing a slice
	fmt.Println(x[1:3])
	fmt.Println(x[2:]) // Till end

	x = append(x, 9, 10)
	fmt.Println(x)
	//x = append(1,2,3,x) --> will not work
	fmt.Println(x)
	y := []int{11, 12, 13, 14}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)

	//deleting from slice
	x = append(x[:2], x[5:]...)
	fmt.Println(x)
}
