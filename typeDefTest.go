package main

func main() {
	type (
		MyInt int
		Age   int
		Text  string
	)

	type IntPtr *int
	type Book struct {
		author, title string
		pages         int
	}
	type Convert func(in0 int, in1 bool) (out0 int, out1 string)
	type StringArray [5]string
	type StringSlice []string

}

func f() {
	type PersonAge map[string]int
	type MessageQueue chan string
	type Reader interface {
		Read([]byte) int
	}
}
