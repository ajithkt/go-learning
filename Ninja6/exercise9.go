package main

import "fmt"

func main() {

	foo(bar)
}

func foo(f func() int) {
	x := f()
	fmt.Println("FOO")
	fmt.Println(x)
}
func bar() int {
	fmt.Println("BAR")
	return 5
}
