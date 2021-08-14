package main

import "fmt"

func main() {

	x := factorial(4)
	fmt.Println(x)
}

func factorial(x int) int {
	total := 1
	for ; x > 0; x-- {
		total = total * x
	}
	return total
}
