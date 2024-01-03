package main

// ReadWriteCloser 此接口直接指定了两个方法和内嵌了两个其他接口
// 其中一个为类型名称，另一个为类型字面量表示形式
type ReadWriteCloser = interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	error                      // 类型名称
	interface{ Close() error } // 类型字面量表示形式
}

// AnyByteSlice 此接口类型内嵌了一个近似类型
// 它的类型集由所有底层类型为[]byte的类型组成。
type AnyByteSlice = interface {
	~[]byte
}

// Unsigned 此接口类型内嵌了一个类型并集。它的类型集包含6个类型：
// uint、uint8、uint16、uint32、uint64、uintptr
type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}
