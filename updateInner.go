//package main
//
//import "fmt"
//
//func Triple(n int) (r int) {
//	defer func() {
//		r += n // 修改返回值
//	}()
//
//	return n + n // <=> r = n + n; return
//}
//
//func main() {
//	fmt.Println(Triple(5)) // 15
//}

package main

import "fmt"

func main() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			i := i
			defer func() {
				fmt.Println("b:", i)
			}()
		}
	}()
}
