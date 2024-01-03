//package main
//
//import "fmt"
//
//func main() {
//	values := []interface{}{
//		456, "abc", true, 0.33, int32(789),
//		[]int{1, 2, 3}, map[int]bool{}, nil,
//	}
//	for _, x := range values {
//		// 虽然v只被声明了一次
//		// 但是它在不同的分支中可以表现为多个类型的变量值
//		switch v := x.(type) {
//		case []int: //一个类型字面量表示
//			// 在此分支中，v的类型为[]int
//			fmt.Print("int slice:", v)
//		case string: // 一个类型名
//			// 在此分支中，v的类型为string
//			fmt.Println("string:", v)
//		case int, float64, int32: // 多个类型名
//			// 在此分支中，v的类型为x的类型interface{}
//			fmt.Println("number:", v)
//		case nil:
//			// 在此分支中，v的类型为x的类型interface{}
//			fmt.Println(v)
//		default:
//			// 在此分支中，v的类型为x的类型interface{}
//			fmt.Println("others:", v)
//		}
//		// 注意：在最后的三个分支中，v均为接口值x的一个复制。
//	}
//}

package main

import "fmt"

func main() {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		if v, ok := x.([]int); ok {
			fmt.Println("int slice:", v)
		} else if v, ok := x.(string); ok {
			fmt.Println("string:", v)
		} else if x == nil {
			v := x
			fmt.Println(v)
		} else {
			_, isInt := x.(int)
			_, isFloat64 := x.(float64)
			_, isInt32 := x.(int32)
			if isInt || isFloat64 || isInt32 {
				v := x
				fmt.Println("number:", v)
			} else {
				v := x
				fmt.Println("others:", v)
			}
		}
	}
}
