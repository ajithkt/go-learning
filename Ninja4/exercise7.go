package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "shaken"}
	y := []string{"Money Penny", "New One"}
	z := [][]string{x, y}
	fmt.Println(z)
	for i, v := range z {
		fmt.Println(i, v)
		for j, v2 := range v {
			fmt.Println(j, v2)
		}

	}
}
