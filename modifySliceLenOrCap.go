package main

import "fmt"

func main() {
	//s := make([]int, 2, 6)
	//fmt.Println(len(s), cap(s))
	//
	//reflect.ValueOf(&s).Elem().SetLen(3)
	//fmt.Println(len(s), cap(s))
	//
	//reflect.ValueOf(&s).Elem().SetCap(5)
	//fmt.Println(len(s), cap(s))
	s := []int{1, 2, 3, 4, 5, 6}
	//from := 2
	//to := 5
	// s = append(s[:from], s[to:]...)
	// fmt.Println(s, len(s), cap(s))

	// copy返回复制了多少个元素
	//s = s[:from+copy(s[from:], s[to:])]
	//fmt.Println(s)

	// 不保持剩余元素的次序
	//if n := to - from; len(s)-to < n {
	//	copy(s[from:], s[to:])
	//} else {
	//	copy(s[from:to], s[len(s)-n:])
	//}
	//s = s[:len(s)-(to-from)]
	//fmt.Println(s)

	i := 3
	//删除第i个元素，保持剩余元素的次序
	//第一种方法
	s = append(s[:i], s[i+1:]...)
	fmt.Println(s)

	//第二种方法
	s = s[:i+copy(s[i:], s[i+1:])]
	fmt.Println(s)

	//第三种方法（不保持剩余元素的次序）
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println(s)
}
