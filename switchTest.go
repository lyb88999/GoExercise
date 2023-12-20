package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//switch n := rand.Intn(100); n % 9 {
	//	//case 0:
	//	//	fmt.Println(n, "is a multiple of 9.")
	//	//
	//	//	// 和很多其它语言不一样，程序不会自动从一个
	//	//	// 分支代码块跳到下一个分支代码块去执行。
	//	//	// 所以，这里不需要一个break语句。
	//	//case 1, 2, 3:
	//	//	fmt.Println(n, "mod 9 is 1, 2 or 3.")
	//	//	break // 这里的break语句可有可无的，效果
	//	//	// 是一样的。执行不会跳到下一个分支。
	//	//case 4, 5, 6:
	//	//	fmt.Println(n, "mod 9 is 4, 5 or 6.")
	//	//// case 6, 7, 8:
	//	//// 上一行可能编译不过，因为6和上一个case中的
	//	//// 6重复了。是否能编译通过取决于具体编译器实现。
	//	//default:
	//	//	fmt.Println(n, "mod 9 is 7 or 8.")
	//	//}

	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		fallthrough // 跳到下个代码块
	case 5, 6, 7, 8:
		// 一个新声明的n，它只在当前分支代码快内可见。
		n := 99
		fmt.Println("n =", n) // 99
		fallthrough
	default:
		// 下一行中的n和第一个分支中的n是同一个变量。
		// 它们均为switch表达式"n"。
		fmt.Println("n =", n)
	}
}
