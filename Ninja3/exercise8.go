package main

import (
	"fmt"
)

func main() {

	//fallthrough will perform all the tasks below even if the condition is true or false or even if it is default condition.
	switch {
	case 2 == 3:
		fmt.Println("Its False")
	case 3 == 3:
		fmt.Println("Its True number 3")
		fallthrough
	case 2 == 2:
		fmt.Println("Its True number 2")
		fallthrough
	case 4 == 2:
		fmt.Println("Its False number 4")
		fallthrough
	case 4 == 4:
		fmt.Println("Its True number 4")
		fallthrough
	case 5 == 5:
		fmt.Println("Its True number 5")
		fallthrough
	default:
		fmt.Println("Its True default")

	}
}
