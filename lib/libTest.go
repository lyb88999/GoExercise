package lib

import "fmt"

func Test1() {
	fmt.Println(123)
}

func init() {
	fmt.Println("libTest111")
}

func init() {
	fmt.Println("libTest222")
}
