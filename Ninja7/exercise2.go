package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changeme(p *person) {
	p.first = "Ajith"
	p.last = "KT"
	p.age = 36
	fmt.Println(p.first, p.last, p.age)
	fmt.Println((*p).first, (*p).last, (*p).age)
}

func main() {

	p1 := person{
		first: "Saranya",
		last:  "B",
		age:   30,
	}

	changeme(&p1)

	fmt.Printf("%T\n", &p1)

}
