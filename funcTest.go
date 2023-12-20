package main

import "fmt"

func squaresOfSumAndDiff(a int64, b int64) (int64, int64) {
	return (a + b) * (a + b), (a - b) * (a - b)
}

func CompareLower4bits(m, n uint32) (r bool) {
	r = m&0xF > n&0xF
	return
}

// 使用一个函数调用的返回结果来初始化一个包级变量
var v = VersionString()

func main() {
	fmt.Println(v)
	x, y := squaresOfSumAndDiff(3, 6)
	fmt.Println(x, y)
	b := CompareLower4bits(uint32(x), uint32(y))
	fmt.Printf("%b\n", 0xF)           // 1111
	fmt.Printf("%b\n", uint32(x))     // 1010001
	fmt.Printf("%b\n", uint32(x)&0xF) //1
	fmt.Printf("%b\n", uint32(y))     //1001
	fmt.Printf("%b\n", uint32(y)&0xF) //1001
	fmt.Println(b)
	doNothing("Go", 1)
}

func VersionString() string {
	return "v1.0"
}

func doNothing(string, int32) {

}
