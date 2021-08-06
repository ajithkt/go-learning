package main

import (
	"fmt"
)

func main() {

	switch {
	case 2 == 3:
		fmt.Println("Its False")
	case 3 == 3:
		fmt.Println("Its True number 3")
	case 2 == 2:
		fmt.Println("Its True number 2")
	}
}
