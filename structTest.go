package main

import "fmt"

type Book struct {
	title, author string
	pages         int
}

func main() {
	book := Book{"Go语言101", "lyb", 256}
	fmt.Println(book)

	book = Book{author: "laoMo", pages: 256, title: "Go101"}
	fmt.Println(book)

	book = Book{}
	fmt.Println(book)

	book = Book{author: "LaoMo12138"}
	fmt.Println(book)

	var book2 Book
	book2.author = "Tapir"
	book2.title = "Go从入门到放弃"
	book2.pages = 300
	fmt.Println(book2)
	fmt.Println(book2.author)
	func() {
		book1 := Book{pages: 300}
		book2 := Book{"Go语言101", "LaoMo", 256}
		book2 = book1
		fmt.Println(book2)
	}()
	book3 := Book{
		title:  "123",
		author: "123",
		pages:  0,
	}
	_ = book3
	fmt.Println(&book3.pages)

	p := &Book{
		title:  "123",
		author: "123",
		pages:  0,
	}
	fmt.Printf("%T %T", book2, p)
}
