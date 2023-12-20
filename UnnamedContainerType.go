package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//// 数组类型
	//const Size = 32
	//[5]string
	//[Size]int
	//[16][]byte
	//[100]Person
	//
	//// 切片类型
	//[]bool
	//[]int64
	//[]map[int]bool
	//[]*int
	//
	////映射类型
	//map[string]int
	//map[int]bool
	//map[int16][6]string
	//map[bool][]string
	//map[struct{x int}]*int8

	// 一个含有4个布尔元素的数组值
	A := [4]bool{false, true, true, false}
	// 一个含有三个字符串值的切片值
	B := []string{"break", "continue", "fallthrough"}
	// 一个映射值
	C := map[string]int{"C": 1972, "Python": 1991, "Go": 2009}

	fmt.Println(A[0])
	fmt.Println(B[2])
	const D string = "Go"
	fmt.Println(C[D])

}
