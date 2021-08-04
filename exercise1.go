package main

import ("fmt")

func main() {
	x := 42
	y := "James, Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Print("x\n", "y\n", "z\n")
	//fmt.Print(x\n, y\n, z\n) --> How to put a new line after a int variable
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
