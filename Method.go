package main

//
//import "fmt"
//
//// Age
//// Age和int是两个不同的类型。我们不能为int和*int
//// 类型声明方法，但是可以为Age和*Age类型声明方法
//type Age int
//
//func (age Age) LargerThan(a Age) bool {
//	return age > a
//}
//func (age *Age) Increase() {
//	*age++
//}
//
//// FilterFunc
//// 为自定义的函数类型FilterFunc声明方法
//type FilterFunc func(in int) bool
//
//func (ff FilterFunc) Filter(in int) bool {
//	return ff(in)
//}
//
//// StringSet
//// 为自定义的映射类型StringSet声明方法
//type StringSet map[string]struct{}
//
//func (ss StringSet) Has(key string) bool {
//	_, present := ss[key]
//	return present
//}
//func (ss StringSet) Add(key string) {
//	ss[key] = struct{}{}
//}
//func (ss StringSet) Remove(key string) {
//	delete(ss, key)
//}
//
//// Book
//// 为自定义的结构体类型Book和它的指针类型*Book声明方法
//type Book struct {
//	pages int
//}
//
//func (b Book) Pages() int {
//	return b.pages
//}
//
//func (b *Book) SetPages(pages int) {
//	b.pages = pages
//}
//
//func main() {
//	var age Age = 3
//	fmt.Println(age)               // 3
//	fmt.Println(age.LargerThan(2)) // true
//	fmt.Println(age.LargerThan(4)) // false
//	age.Increase()
//	fmt.Println(age) // 4
//
//	myFilter := FilterFunc(func(in int) bool {
//		return in > 5
//	})
//	fmt.Println(myFilter(3))        //false
//	fmt.Println(myFilter.Filter(6)) //true
//
//	set := StringSet{
//		"XiaoLi":   {},
//		"Xiaoming": {},
//	}
//
//	fmt.Println("\n")
//	fmt.Println(set.Has("XiaoLi")) // true
//	fmt.Println(set.Has("Tom"))    // false
//	set.Add("Tom")
//	fmt.Println(set.Has("Tom")) // true
//	set.Remove("Tom")
//	fmt.Println(set.Has("Tom")) // false
//	fmt.Println(set)
//
//	book := Book{
//		400,
//	}
//	fmt.Println(book.pages) // 400
//	book.SetPages(132)
//	fmt.Println(book.pages) // 132
//}
