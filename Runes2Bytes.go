package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func Runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}

func main() {
	s := "颜色感染是一个有趣的游戏。"
	bs := []byte(s) // string -> []byte
	fmt.Println(bs)
	s = string(bs) // []byte -> string
	fmt.Println(s)
	rs := []rune(s) // string -> []rune
	fmt.Println(rs)
	s = string(rs) // []rune -> string
	fmt.Println(s)
	rs = bytes.Runes(bs) // []byte -> []rune
	fmt.Println(rs)
	bs = Runes2Bytes(rs) // []rune -> []byte
	fmt.Println(bs)
}
