package main

import (
	"fmt"
)

// Raw string liternals using ``
func main() {
	x := `Hello, %d I am good
		"How are YOU %u`
	fmt.Println(x)
}
