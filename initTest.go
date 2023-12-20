package main

import "fmt"
import "test/lib"

func init() {
	fmt.Println("hi,", bob)
}

func main() {
	fmt.Println("bye")
	lib.Test1()
}

func init() {
	fmt.Println("hello,", smith)
}

func titledName(who string) string {
	return "Mr. " + who
}

var bob, smith = titledName("Bob"), titledName("Smith")
