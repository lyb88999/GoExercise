package main

import "fmt"

func HalfAndNegative(n int) (int, int) {
	return n / 2, -n
}

func AddSub(a, b int) (int, int) {
	return a + b, a - b
}

func Dummy(values ...int) {}

func main() {
	fmt.Println(AddSub(HalfAndNegative(6)))         // AddSub(3,-6) -> -3 9
	fmt.Println(AddSub(AddSub((AddSub(7, 5)))))     // AddSub(AddSub(12,2)) -> AddSub(14,10) -> 24,4
	fmt.Println(AddSub(AddSub(HalfAndNegative(6)))) // AddSub(AddSub(3,-6)) -> AddSub(-3,9) -> 6,-12
	Dummy(HalfAndNegative(6))
	_, _ = AddSub(7, 5)

	// 下面这几行编译不通过。
	/*
		_, _, _ = 6, AddSub(7, 5)
		Dummy(AddSub(7, 5), 9)
		Dummy(AddSub(7, 5), HalfAndNegative(6))
	*/
}
