package main

import (
	"fmt"
)

type aji int

var x aji

func main() {
	fmt.Println(x)
	fmt.Printf("%T	1\n", x)
	x = 42
	fmt.Println(x)
}
