package main

import (
	"fmt"
)

func main() {
	var i interface{}
	i = []int{1, 2, 3}
	fmt.Println(i)
	i = map[string]int{"Go": 2012}
	fmt.Println(i)
	i = true
	fmt.Println(i)
	i = 1
	fmt.Println(i)
	i = "abc"
	fmt.Println(i)

	//将接口值i中包裹的值清除掉
	i = nil
	fmt.Println(i)
}
