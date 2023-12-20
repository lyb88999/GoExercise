package main

import "fmt"

func main() {
	//fmt.Println(0b1100 & 0b1010)
	//fmt.Println(0b1100 | 0b1010)
	//fmt.Println(0b1100 ^ 0b1010)
	//fmt.Println(0b1100 >> 1) // 0b110->6
	//fmt.Println(0b1100 << 1) // 0b11000->24
	//fmt.Println(0b1110 >> 2) //0b11->3
	//const a int8 = -128
	//fmt.Println(a >> 2) //10000000->11100000->-32
	//const b int8 = -128
	//fmt.Printf("%d", b>>2)
	var a, b uint8 = 255, 1
	var c = a + b
	fmt.Printf("%b", c)
	var d = a << b
	fmt.Printf("%b\n", a)
	fmt.Printf("%b", d)
	const X = 0x1FFFFFFFF * 0x1FFFFFFFF
	const R = 'a' + 0x7FFFFFFFF

}
