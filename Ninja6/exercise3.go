package main

import (
	"fmt"
)

func main() {
	defer foo() //This will run only at last when the outside function( func which contain this function) exits.
	bar()
	fmt.Println("Hello world")
}

func foo() {
	defer func() { fmt.Println("HELLO FROM ANONYMOUS FUNCTION") }() //Anonymous function and it execution using () at last
	fmt.Println("Hello world from foo")
}

func bar() {
	fmt.Println("Hello world from bar")
}
