package main

import "fmt"

type S []int
type A [2]int
type P *A

func main() {
	//var x []int
	//var y = make([]int, 0)
	//var x0 = (*[0]int)(x) // okay, y0 == nil
	//var y0 = (*[0]int)(y) // okay, y0 != nil
	//_, _ = x0, y0
	//
	//var z = make([]int, 3, 5)
	//z[0], z[1], z[2] = 3, 4, 5
	//var _ = (*[3]int)(z) // okay
	//var _ = (*[2]int)(z) // okay
	//test := (*[2]int)(z) // okay
	//_ = test
	//var _ = (*A)(z) // okay
	//var _ = P(z)    // okay
	//
	//var w = S(z)
	//var _ = (*[3]int)(w) // okay
	//var _ = (*[2]int)(w) // okay
	//var _ = (*A)(w)      // okay
	//var _ = P(w)         // okay
	//
	//var _ = (*[4]int)(z) // 会产生恐慌
	var s = []int{0, 1, 2, 3}
	var a = [3]int(s[1:])
	s[2] = 9
	fmt.Println(s) // [0 1 9 3]
	fmt.Println(a) // [1 2 3]

	_ = [3]int(s[:2]) // 会产生恐慌
}
