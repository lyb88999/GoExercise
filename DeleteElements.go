package main

//
//import "fmt"
//
//type T int
//
//func BiggerThan5(t T) bool {
//	if t > 5 {
//		return true
//	} else {
//		return false
//	}
//}
//
//func DeleteElements(s []T, keep func(T) bool, clear bool) []T {
//	result := s[:0]
//	for _, v := range s {
//		if keep(v) {
//			result = append(result, v)
//		}
//	}
//	if clear {
//		temp := s[len(result):]
//		for i := range temp {
//			temp[i] = 0
//		}
//	}
//	return result
//}
//func main() {
//	// s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
//	res := DeleteElements([]T{0, 1, 2, 3, 4, 5, 6, 7, 8}, BiggerThan5, true)
//	fmt.Println(res) // [6 7 8]
//}
