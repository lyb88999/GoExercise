package main

import "fmt"

func main() {
	var c chan struct{} // nil
	select {
	case <-c: // 阻塞操作 尝试从通道c接收数据。但由于c是一个nil通道，这个操作会一直阻塞。
	case c <- struct{}{}: // 阻塞操作 尝试向通道c发送一个空结构体。同样，由于c是nil，这个操作也会一直阻塞。
	default:
		fmt.Println("Go here.")
	}
}

// 由于通道c是nil，它既不能接收也不能发送数据，所以两个case
// 都会阻塞。因此，程序会执行default情况并打印"Go here."。
// 这个程序演示了如何使用select来处理可能的阻塞操作，并提供了一种不阻塞的默认操作。
