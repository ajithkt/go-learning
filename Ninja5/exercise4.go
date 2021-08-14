package main
map[string]int

import (
	"fmt"
)

//anonymous struct
func main() {
	x := struct {
		first  string
		last   string
		age    int
		marks  map[string]int
		grades []string
	}{
		first: "Ajith",
		last:  "KT",
		age:   36,
		marks: map[string]int{
			"maths":    90,
			"science":  80,
			"computer": 95,
		},
		grades: []string{"A", "B", "A"},
	}
	fmt.Println(x)
	fmt.Println(x.grades[1])
}
