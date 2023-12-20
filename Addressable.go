package main

import "fmt"

func main() {
	//a := [5]int{2, 3, 5, 7}
	//s := make([]bool, 2)
	//pa2, ps1 := &a[2], &s[1]
	//fmt.Println(*pa2, *ps1) // 5 false
	//a[2], s[1] = 99, true
	//fmt.Println(*pa2, *ps1) // 99 true
	//ps0 := &[]string{"Go", "C"}[0]
	//fmt.Println(*ps0) // Go
	//
	//m := map[int]bool{1: true}
	//_ = m
	//
	//// 下面几行编译不通过
	///*
	//	_ = &[3]int{2, 3, 5}[0]
	//	_ = &map[int]bool{1: true}[1]
	//	_ = &m[1]
	//*/

	type T struct{ age int }
	mt := map[string]T{}
	mt["John"] = T{age: 29}
	ma := map[int][5]int{}
	ma[1] = [5]int{1: 789}

	fmt.Println(ma[1][1])
	fmt.Println(mt["John"].age)
}
