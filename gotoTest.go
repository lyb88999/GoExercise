package main

import "fmt"

func main() {
	i := 0
Next:
	fmt.Println(i)
	i += 1
	if i > 10 {
		return
	}
	goto Next
}
