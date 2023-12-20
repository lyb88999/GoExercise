package main

import "fmt"

func FindSmallestPrimeLargerThan(n int) int {
Outer:
	for n++; ; n++ {
		for i := 2; ; i++ {
			switch {
			case i*i > n:
				break Outer
			case n%i == 0:
				continue Outer
			}
		}
	}
	return n
}

func main() {
	for i := 0; i < 100; i++ {
		n := FindSmallestPrimeLargerThan(i)
		fmt.Print("最小的比", i, "大的素数为", n)
		fmt.Println()
	}
}
