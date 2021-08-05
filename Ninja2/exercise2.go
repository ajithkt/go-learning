package main

import (
	"fmt"
)

func main() {

	x := 42
	y := 52
	z := (x == y)
	fmt.Println(z)
	z = (x <= y)
	fmt.Println(z)
	z = (x >= y)
	fmt.Println(z)
	z = (x < y)
	fmt.Println(z)
	z = (x > y)
	fmt.Println(z)
}
