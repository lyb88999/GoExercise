package main

import "fmt"

func double(x *int) {
	*x += *x
}
func main() {
	//var a int64 = 3
	//var b int64 = 4
	//var c = new(int64)
	//fmt.Println(&a, &b, &c)
	x := 3
	double(&x)
	fmt.Println(x)
}
