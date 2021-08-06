package main

import (
	"fmt"
)

func main() {
	i := 0
	y := 1985
	for {

		if i > 36 {
			break
		}
		fmt.Println(y + i)
		i++

	}
}
