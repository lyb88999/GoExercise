package main

import "time"

type Buffer struct {
	start, end int
	data       [1024]byte
}

func fa(buffers []Buffer) int {
	numUnreads := 0
	for _, buf := range buffers {
		numUnreads += buf.end - buf.start
	}
	return numUnreads
}

func fb(buffers []Buffer) int {
	numUnreads := 0
	for i := range buffers {
		numUnreads += buffers[i].end - buffers[i].start
	}
	return numUnreads
}

func main() {
	buffers := []Buffer{
		0: {start: 0, end: 512, data: [1024]byte{1}},
		1: {start: 0, end: 1024, data: [1024]byte{2}},
		2: {start: 0, end: 512, data: [1024]byte{3}},
	}
	// 计算运行时间
	start := time.Now()
	fa(buffers)
	end := time.Now()
	println(end.Sub(start).String()) // 111ns
	start = time.Now()
	fb(buffers)
	end = time.Now()
	println(end.Sub(start).String()) // 103ns
}
