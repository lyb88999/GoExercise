package main

import (
	"fmt"
	"strings"
)

func main() {
	var helloWorld = "hello world!" // 取子字符串
	var hello = helloWorld[:5]
	// 104是英文字符h的ASCII（和Unicode）码
	fmt.Println(hello[0])         // 104
	fmt.Printf("%T \n", hello[0]) //unit8

	// hello[0]是不可寻址和不可修改的
	// 下面两行编译不通过
	/*
		hello[0] = 'H' // error
		fmt.Println(&hello[0]) // error
	*/
	// 下一条语句将打印出：5 12 true
	fmt.Println(len(hello), len(helloWorld))
	fmt.Println(strings.HasPrefix(helloWorld, hello))
}
