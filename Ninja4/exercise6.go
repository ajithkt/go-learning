package main

import (
	"fmt"
)

func main() {

	x := make([]string, 5, 5)
	fmt.Println(x)
	x = []string{"a", "b", "c", "d"}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, "e", "f", "g")
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Print(x[i])
	}
}
