package main

import "math/rand"

// MaxRand 是具名常量 16 是整形字面量
const MaxRand = 16

// 一个函数声明
/*
 StatRandomNumbers生成一些不大于MaxRand的非负
 随机整数，并统计和返回小于和大于MaxRand/2的随机数
 个数。输入参数numRands指定了要生成的随机数的总数。
*/

func StatRandomNumbers(numRands int) (int, int) {
	// 声明了两个变量(类型都为int, 初始值都为0)
	var a, b int
	// 一个for循环代码块
	// 大括号不能放在下一行 否则会编译错误
	for i := 0; i < numRands; i++ {
		if rand.Intn(MaxRand) < MaxRand/2 {
			a = a + 1
		} else {
			b++
		}
	}
	return a, b
}

func main() {
	var num = 100
	x, y := StatRandomNumbers(num)
	// "Result: "是一个字符串字面量
	print("Result: ", x, "+", y, "=", num, "? ")
	println(x+y == num)
}
