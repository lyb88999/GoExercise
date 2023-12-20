package main

import "fmt"

func main() {
	//const (
	//	Failed    = iota - 1 // == -1
	//	Unknown              // == 0
	//	Succeeded            // == 1
	//)
	//
	//const (
	//	Readable   = 1 << iota // == 1
	//	Writable               // == 2
	//	Executable             // == 4
	//)
	//fmt.Println(Failed, Unknown, Succeeded)
	//fmt.Println('\n')
	//fmt.Println(Readable, Writable, Executable)
	const N = 123
	var x int
	var y, z float32
	// N = 789 error
	y = N
	x = int(y)
	// x = y  error
	x = N
	// y = x error
	z = y
	_ = y
	z, y = y, z
	_, y = y, z
	z, _ = y, z
	_, _ = y, z
	x, y = 69, 1.23
	// x, y = y, x error 类型不匹配
	x, y = int(y), float32(x)
	fmt.Println(x, y)
}
