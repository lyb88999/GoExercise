package main

//import "fmt"
//
//type Aboutable interface {
//	About() string
//}
//
//type Book struct {
//	name string
//}
//
//// About 类型*Book实现了接口类型Aboutable。
//func (book *Book) About() string {
//	return "Book: " + book.name
//}
//
//func main() {
//	// 一个*Book值被包裹在了一个Aboutable值中。
//	// 创建了一个Book的实例，并将其地址（即指针）赋给了Aboutable类型的变量a。
//	// 因为*Book实现了Aboutable接口，所以这是允许的。
//	var a Aboutable = &Book{"Go语言101"}
//	fmt.Println(a) // &{Go语言101}
//
//	// i是一个空接口值。类型*Book实现了任何空接口类型。
//	// 创建了另一个Book实例，并将其地址赋给了一个空接口变量i。
//	// 空接口（interface{}）可以持有任何类型的值。
//	var i interface{} = &Book{"Rust 101"}
//	fmt.Println(i) // &{Rust 101}
//
//	// Aboutable实现了空接口类型interface{}。
//	// 由于i是一个空接口，它可以接受任何类型的值，包括Aboutable接口类型的值。
//	i = a
//	fmt.Println(i) // &{Go语言101}
//}
