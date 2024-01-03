package main

type MyInt int

func (mi MyInt) IsOdd() bool {
	return mi%2 == 1
}

type Age MyInt

func main() {
	var x MyInt = 3
	_ = x.IsOdd()
	var y Age = 36
	// _ = y.IsOdd() // error: y.IsOdd undefined
	_ = y
}
