package main

import "fmt"

func main() {
	for i := 33; i < 123; i++ {
		// Printing ascii code
		fmt.Printf("%d \t %q \t %#U\n", i, i, i)
	}
}
