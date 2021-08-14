package main

import (
	"fmt"
)

func main() {

	v := 10
	fmt.Println("Value is:", v)
	fmt.Println("Address is:", &v)
	fmt.Println("De-referencing address to value is:", *&v)
}
