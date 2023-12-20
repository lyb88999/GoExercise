package main

import "fmt"

func main() {
	x, y := func() (int, int) {
		fmt.Println("This function has no parameters.")
		return 3, 4
	}()

	func(a, b int) {
		fmt.Println("a*a+b*b=", a*a+b*b)
	}(x, y)

	func(x int) {
		fmt.Println("x*x+y*y=", x*x+y*y)
	}(y)

	func() {
		fmt.Println("x*x+y*y=", x*x+y*y)
	}()
}
