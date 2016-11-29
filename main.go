package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")

	a := aa()

	fmt.Println(a)
}

func aa() string {
	return "asd"
}
