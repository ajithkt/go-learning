package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "Ajith",
		last:  "KT",
		age:   36,
	}
	p2 := person{
		first: "Saranya",
		last:  "B",
		age:   31,
	}
	p1.speak()
	p2.speak()
}

func (p person) speak() {
	fmt.Println("Iam", p.first, " ", p.last, " & My age is ", p.age)
}
