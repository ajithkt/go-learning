package main

import "fmt"

func main() {

	fmt.Printf("%v\tAND\t%v\t%v\n", true, true, true && true)
	fmt.Printf("%v\tAND\t%v\t%v\n", true, false, true && false)
	fmt.Printf("%v\tOR\t%v\t%v\n", true, true, true || true)
	fmt.Printf("%v\tOR\t%v\t%v\n", true, false, true || false)
	fmt.Printf("%v\tOpposite is\t%v", true, !true)
}
