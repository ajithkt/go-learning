package main

import "fmt"

func main() {

	x := foo()
	x()
	fmt.Println("%T", x)

}

func foo() func() {
	return func() { //return a function
		fmt.Println("You are inside bar")
	}
}
