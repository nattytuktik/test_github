package main

import (
	"fmt"
)

func main() {
	var i interface{} = "hello"

	s, ok := i.(string)

	f, okf := i.(float32)

	fmt.Println(s, ok)

	fmt.Println(f, okf)
}
