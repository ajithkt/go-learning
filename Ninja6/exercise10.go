package main

import "fmt"

var x int = 5

func main() {

	fmt.Println(x)
	x = 42
	fmt.Println(x)
	{
		var y string = "Closure variable"
		fmt.Println(x)
		fmt.Println(y)
	}
	//fmt.Println(y) --> closure variable scope is confined.
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println("-----")
	g := foo()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func foo() func() int {
	z := 0
	return func() int {
		z++
		return z
	}
}
