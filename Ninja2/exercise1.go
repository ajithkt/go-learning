package main

import (
	"fmt"
)

var x = 42

func main() {
	//Decimal
	fmt.Printf("%d\n", x)
	//Binary
	fmt.Printf("%b\n", x)
	//Hexa-decimal
	fmt.Printf("%x\n", x)
	//Hexadecimal with 0x format.
	fmt.Printf("%#x", x)

}
