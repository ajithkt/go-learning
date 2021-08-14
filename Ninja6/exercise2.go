package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum1 := foo(x...)
	sum2 := bar(x)
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func foo(s ...int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

func bar(s []int) int {
	total2 := 0
	for _, v := range s {
		total2 += v
	}
	return total2
}
