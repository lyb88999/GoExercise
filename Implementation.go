package main

//import (
//	"database/sql"
//	"fmt"
//)
//
//type Aboutable interface {
//	About() string
//}
//
//type Book struct {
//	name string
//}
//
//func (book *Book) About() string {
//	return "Book: " + book.name
//}
//
//type MYInt int
//
//func (MYInt) About() string {
//	return "我是一个自定义整数类型"
//}
//
//type DatabaseStorer interface {
//	Exec(query string, args ...interface{}) (sql.Result, error)
//	Prepare(query string) (*sql.Stmt, error)
//	Query(query string, args ...interface{}) (*sql.Rows, error)
//}
//
//func main() {
//	book := Book{
//		"我和我的祖国",
//	}
//	fmt.Println(book.About())
//
//	myInt := MYInt(3)
//	fmt.Println(myInt.About())
//}
