package main

import "fmt"

func main() {
	m := map[string]int{"C": 1972, "C++": 1983, "Go": 2009}
	for lang, year := range m {
		fmt.Printf("%v: %v\n", lang, year)
	}

	a := [...]int{2, 3, 5, 7, 11}
	for idx, num := range a {
		fmt.Printf("%v: %v\n", idx, num)
	}

	s := []string{"go", "defer", "goto", "var"}
	for idx, keyword := range s {
		fmt.Printf("%v: %v\n", idx, keyword)
	}
}
