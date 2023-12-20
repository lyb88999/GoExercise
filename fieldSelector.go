package main

type S0 struct {
	y int `foo:"bar"`
	x bool
}

type S1 = struct { //S1是一个无名类型
	x int `foo:"bar"`
	y bool
}

type S2 = struct { //S2也是一个无名类型
	x int `bar:"foo"`
	y bool
}

type S3 S2 // S3是一个定义类型
type S4 S3 // S4是一个定义类型
// 如果不考虑字段标签，S3(S4)和S1的底层类型一样
// 如果考虑字段标签，S3(S4)和S1的底层类型不一样

var v0, v1, v2, v3, v4 = S0{}, S1{}, S2{}, S3{}, S4{}

func main() {
	v1 = S1(v2)
	v2 = S2(v1) // S1和S2的底层类型相同（忽略掉字段标签） 可以相互转换为对方类型
	v1 = S1(v3)
	v3 = S3(v1) // S1和S3的底层类型相同（忽略掉字段标签） 可以相互转换为对方类型
	v1 = S1(v4)
	v4 = S4(v1) // S1和S4的底层类型相同（忽略掉字段标签） 可以相互转换为对方类型
	v2 = v3
	v3 = v2 // S2和S3的底层类型相同（考虑字段标签）
	v2 = v4
	v4 = v2     // S2和S4的底层类型相同（考虑字段标签)且S2为无名类型，可以隐式转换
	v3 = S3(v4) // S3和S4底层类型相同（考虑字段标签)但是都是具名类型，不能隐式转换
	v4 = S4(v3)
	//type Book struct {
	//	pages int
	//}
	//book1 := &Book{100}
	//book2 := new(Book)
	//book2.pages = book1.pages
	//fmt.Println(*book1)
	//fmt.Println(*book2)
	//fmt.Println(*book1 == *book2)
}
