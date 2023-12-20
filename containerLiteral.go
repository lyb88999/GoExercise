package main

import "fmt"

func main() {
	//pm := &map[string]int{"C": 1972, "Go": 2009}
	//ps := &[]string{"break", "continue"}
	//pa := &[...]bool{false, true, true, false}
	//fmt.Printf("%T\n", pm)
	//fmt.Printf("%T\n", ps)
	//fmt.Printf("%T\n", pa)
	//
	//var heads = []*[4]byte{
	//	&[4]byte{'P', 'N', 'G', ' '},
	//	&[4]byte{'G', 'I', 'F', ' '},
	//	&[4]byte{'J', 'P', 'E', 'G'},
	//}
	//_ = heads
	//var headsSimplified = []*[4]byte{
	//	{'P', 'N', 'G', ' '},
	//	{'G', 'I', 'F', ' '},
	//	{'J', 'P', 'E', 'G'},
	//}
	//_ = headsSimplified
	//
	//type language struct {
	//	name string
	//	year int
	//}
	//
	//var _ = [...]language{
	//	language{"C", 1972},
	//	language{"Python", 1991},
	//	language{"Go", 2009},
	//}
	//
	//var _ = [...]language{
	//	{"C", 1972},
	//	{"Python", 1991},
	//	{"Go", 2009},
	//}
	//
	//type LangCategory struct {
	//	dynamic bool
	//	strong  bool
	//}
	//
	//var _ = map[LangCategory]map[string]int{
	//	LangCategory{true, true}: map[string]int{
	//		"Python": 1991,
	//		"Erlang": 1986,
	//	},
	//	LangCategory{true, false}: map[string]int{
	//		"JavaScript": 1995,
	//	},
	//	LangCategory{false, true}: map[string]int{
	//		"Go":   2009,
	//		"Rust": 2010,
	//	},
	//	LangCategory{false, false}: map[string]int{
	//		"C": 1972,
	//	},
	//}
	//
	//var _ = map[LangCategory]map[string]int{
	//	{true, true}: {
	//		"Python": 1991,
	//		"Erlang": 1986,
	//	},
	//	{true, false}: {
	//		"JavaScript": 1995,
	//	},
	//	{false, true}: {
	//		"Go":   2009,
	//		"Rust": 2010,
	//	},
	//	{false, false}: {
	//		"C": 1972,
	//	},
	//}
	//
	var a [16]byte
	var s []int
	var m map[string]int

	fmt.Println(a == a)
	fmt.Println(s == nil)
	fmt.Println(m == nil)
	fmt.Println(nil == map[string]int{})
	fmt.Println(nil == []int{})

	// 下面这些行编译不通过。
	/*
		_ = m == m
		_ = s == s
		_ = m == map[string]int(nil)
		_ = s == []int(nil)
		var x [16][]int
		_ = x == x
		var y [16]map[int]bool
		_ = y == y
	*/
}
