package main

import (
	"fmt"
)

func main() {

	var m map[string]string

	m = map[string]string{
		"ajith":   "kt",
		"saranya": "b",
		"Ishaan":  "Ajith",
	}

	fmt.Println(m)

	for i, v := range m {

		fmt.Println(i, v)
	}

	cm := map[string][]string{
		"ajith":   []string{"kt", "36"},
		"saranya": []string{"b", "30"},
		"Ishaan":  []string{"Ajith", "1"},
	}
	fmt.Println(cm)
	for i, v := range cm {
		fmt.Println(i, v)
		for j, v2 := range v {
			fmt.Println(j, v2)
		}
	}

}
