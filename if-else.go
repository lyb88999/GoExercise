package main

import (
	"fmt"
	"time"
)

func main() {
	if h := time.Now().Hour(); h < 12 {
		fmt.Println("现在是上午")
	} else if h > 19 {
		fmt.Println("现在是晚上")
	} else {
		fmt.Println("现在是下午")
		h := h
		_ = h
	}
}
