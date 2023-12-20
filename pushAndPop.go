package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	var e int
	// 前弹出(pop front/shift)
	s, e = s[1:], s[0]
	// 后弹出(pop back)
	fmt.Println(s, e)
	s, e = s[:len(s)-1], s[len(s)-1]
	fmt.Println(s, e)
	// 前推(push front)
	s = append([]int{0}, s[:]...)
	fmt.Println(s)
	// 后推(push back)
	s = append(s, 6)
	fmt.Println(s)
}
