package main

import "fmt"

func main() {
	const N = 2
	// 如果左操作数是一个类型不确定值并且右操作数是一个常量，则左操作数将总是被视为一个整数。
	const A = 3.0 << N
	const B = int8(3.0) << N
	var m = uint(32)
	var x int64 = 1 << m
	var y = int64(1 << m)
	var z = int64(1) << m
	var l float32 = 1.23
	var r = int(l)
	// var _ = int(1.23)
	fmt.Printf("%b\n", 3.0<<N)
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
	fmt.Printf("%b\n", r)
}
