package main

import "fmt"

const pi = 3.1416
const Pi = pi
const (
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "弧度"
)

func main() {
	// var x = complex128(1 + -1e-1000i)
	// var x = float32(0.49999999)
	// var x = float32(1700000000000000)
	// var x = float32(123)
	// var x = uint(1.0)
	// var x = string(-1)
	//const DoublePi, HalfPi, Unit2 = pi * 2, pi * 0.5, "度"
	// const MaxUint = ^uint(0)
	const MaxUint uint = (1 << 64) - 1
	const MaxInt = int(^uint(0) >> 1)
	const NativeWordBits = 32 << (^uint(0) >> 63)
	const Is64bitOS = ^uint(0)>>63 != 0
	const Is32bitOS = ^uint(0)>>32 == 0
	// fmt.Print(MaxUint, MaxInt)
	fmt.Print(NativeWordBits, Is64bitOS, Is32bitOS)
}
