package main

import (
	"fmt"
	"time"
)

// 边执行边从通道取数据
func main() {
	fibonacci := func() chan uint64 {
		c := make(chan uint64)
		go func() {
			var x, y uint64 = 0, 1
			for ; y < (1 << 63); c <- y { // 步尾语句
				x, y = y, x+y
			}
			close(c)
		}()
		return c
	}
	c := fibonacci()
	//for x, ok := <-c; ok; x, ok = <-c { // 初始化和步尾语句
	//	time.Sleep(time.Second)
	//	fmt.Println(x)
	//}
	for x := range c {
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}
