// package main
//
// import "fmt"
//
//	func main() {
//		var a, b, c interface{} = "abc", 123, "a" + "b" + "c"
//		fmt.Println(a == b) // 动态类型不一样 false
//		fmt.Println(a == c) // 动态类型一样且动态值一样 true
//
//		var x *int = nil
//		var y *bool = nil
//		var ix, iy interface{} = x, y
//		var i interface{} = nil
//		fmt.Println(ix == iy) // 动态类型不一样 false
//		fmt.Println(ix == i)  // 一个接口值是 nil 接口值，另一个不是 false
//		fmt.Println(iy == i)  // 一个接口值是 nil 接口值，另一个不是 false
//
//		var s []int = nil // []int为一个不可比较类型
//		i = s
//		fmt.Println(i == nil) // 一个接口值是 nil 接口值，另一个不是 false
//		fmt.Println(i == i)   // 动态类型一样 且都为不可比较类型 产生一个恐慌
//	}
package main

import "fmt"

var dataSlice []interface{}

func main() {
	dataSlice = append(dataSlice, 42, "hello", struct{}{})
	fmt.Println(dataSlice, len(dataSlice))
}
