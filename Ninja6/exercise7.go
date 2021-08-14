package main

import "fmt"

var g func() = func() { fmt.Println("FUNCTION OUTSIDE MAIN FUNC") }

func main() {
	x := sample
	x()
	g()
	fmt.Printf("%T", x)
}

func sample() {
	fmt.Println("THIS IS A FUNC ASSIGNED TO A VARIABLE")
}
