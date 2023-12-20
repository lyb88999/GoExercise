package main

import "fmt"

func main() {
	fmt.Println(5/3, 5%3)
	fmt.Println(5/-3, 5%-3)
	fmt.Println(-5/3, -5%3)
	fmt.Println(-5/-3, -5%-3)

	fmt.Println(5.0 / 3.0)
	fmt.Println((1 - 1i) / (1 + 1i))
	var a, b = 1.0, 0.0
	fmt.Println(a/b, b/b)
	_ = int(a) / int(b)
	_ = int(a) / int(b)
}
