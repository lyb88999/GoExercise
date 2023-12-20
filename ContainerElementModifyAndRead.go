package main

import "fmt"

func main() {
	a := [3]int{-1, 0, 1}
	s := []bool{true, false}
	m := map[string]int{"abc": 123, "xyz": 789}
	fmt.Println(a[2], s[1], m["abc"])     // 读取
	a[2], s[1], m["abc"] = 999, true, 567 // 修改
	fmt.Println(a[2], s[1], m["abc"])     // 读取

	n, present := m["hello"]
	fmt.Println(n, present, m["hello"]) // 0 false 0
	n, present = m["abc"]
	fmt.Println(n, present, m["abc"]) // 567 true 567
	m = nil
	fmt.Println(m["abc"]) // 0
}
