package main

import "fmt"

// Sum 返回所有输入实参的和。
func Sum(values ...int64) (sum int64) {
	sum = 0
	for _, value := range values {
		sum += value
	}
	return
}

func Concat(sep string, tokens ...string) string {
	r := ""
	for i, token := range tokens {
		if i != 0 {
			r += sep
		}
		r += token
	}
	return r
}
func main() {
	sum := Sum(1, 2, 3, 4, 5)
	fmt.Println(sum)                                     // 15
	r := Concat(",", "Hello", "What are you doing now?") // Hello,What are you doing now?
	fmt.Println(r)
}
