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

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	fmt.Println(m)

	for _, v := range m {

		fmt.Println(v.first_name)
		fmt.Println(v.last_name)
		for i, v2 := range v.ice_cream {
			fmt.Println(i, v2)
		}

	}
}
