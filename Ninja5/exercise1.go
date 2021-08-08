package main

import (
	"fmt"
)

type person struct {
	first_name string
	last_name  string
	ice_cream  []string
}

func main() {

	p1 := person{
		first_name: "Ajith",
		last_name:  "KT",
		ice_cream:  []string{"Vannilla", "Pista"},
	}
	p2 := person{
		first_name: "saranya",
		last_name:  "B",
		ice_cream:  []string{"Chocolate", "Strawberry"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	for i, v := range p1.ice_cream {
		fmt.Println(i, v)
	}

}
