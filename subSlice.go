package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	s0 := a[:]
	s1 := s0[:]
	s2 := s1[1:3]
	s3 := s1[3:]
	s4 := s0[3:5]
	s5 := s4[:2:2]
	s6 := append(s4, 77)
	s7 := append(s5, 88)
	s8 := append(s7, 66)
	s3[1] = 99
	fmt.Println(len(s2), cap(s2), s2)
	fmt.Println(len(s3), cap(s3), s3)
	fmt.Println(len(s4), cap(s4), s4)
	fmt.Println(len(s5), cap(s5), s5)
	fmt.Println(len(s6), cap(s6), s6)
	fmt.Println(len(s7), cap(s7), s7)
	fmt.Println(len(s8), cap(s8), s8)
}
