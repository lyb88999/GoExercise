package main

import "fmt"

func Sum2(values ...int64) (sum int64) {
	sum = 0
	for _, value := range values {
		sum += value
	}
	return
}

func Concat2(sep string, tokens ...string) string {
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
	a0 := Sum2()
	a1 := Sum2(2)
	a3 := Sum2(2, 3, 5)

	// 上面三行和下面三行是等价的
	b0 := Sum2([]int64{}...)
	b1 := Sum2([]int64{2}...)
	b3 := Sum2([]int64{2, 3, 5}...)
	fmt.Println(a0, a1, a3)
	fmt.Println(b0, b1, b3)

	tokens := []string{"Go", "C", "Rust"}
	langA := Concat2(",", tokens...)
	langB := Concat2(",", "Go", "C", "Rust")
	fmt.Println(langA, langB)
	fmt.Println(langA == langB)
}
