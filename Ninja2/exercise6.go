package main

import (
	"fmt"
)

//iota for recuring values will work only for consts
const (
	i = iota
	a
	b
	c
	d
)

func main() {

	fmt.Println(i, "\t", a, "\t", b, "\t", c, "\t", d)

}
