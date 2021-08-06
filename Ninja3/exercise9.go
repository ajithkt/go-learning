package main

import (
	"fmt"
)

func main() {
	x := "favSport"
	switch x {

	case "fav":
		fmt.Println("FAV")
	case "favSport":
		fmt.Println("FAVSPORT")
	case "fav111":
		fmt.Println("FAV111")
	default:
		fmt.Println("SOMETHING DEFAULT")
	}
}
