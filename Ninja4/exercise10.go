package main

import "fmt"

func main() {

	m := map[string][]string{
		"ajith":   []string{"kt", "paloor"},
		"Saranya": []string{"b", "Guruvayoor"},
		"Ishaan":  []string{"Ajith", "Alexandria"},
	}
	m["New one"] = []string{"Awaiting", "US"}
	fmt.Println(m)

	delete(m, "ajith")

	for i, v := range m {
		fmt.Println(i, v)
		for j, v2 := range v {
			fmt.Println(j, v2)
		}
	}
}
