package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	clear(s)
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 3

	a := [4]int{5, 6, 7, 8}
	clear(a[1:3])
	fmt.Println(a, len(a), cap(a)) // [5 0 0 8] 4 4

	m := map[float64]float64{}
	x := 0.0
	m[x] = x
	x /= x
	m[x] = x
	fmt.Println(m, len(m)) // map[NaN:NaN 0:0] 2
	for k := range m {
		delete(m, k)
	}
	fmt.Println(m, len(m)) // map[NaN:NaN] 1
	clear(m)
	fmt.Println(m, len(m)) // map[] 0
}
