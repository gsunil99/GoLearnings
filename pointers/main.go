package main

import (
	"fmt"
)

func main() {
	name := "sunil"
	ptr := &name
	fmt.Println(*ptr)
	fmt.Println(&ptr)
	fmt.Println(ptr)
}
