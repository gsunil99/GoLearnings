package main

import (
	"errors"
	"fmt"
)

func IsEven(num int) error {
	if num%2 != 0 {
		return errors.New("number is not even")
	}
	return nil
}

func main() {
	err := IsEven(24)
	if err != nil {
		fmt.Println("handle the error")
		fmt.Println(err)
	}

	err = IsEven(25)
	if err != nil {
		fmt.Println("handle the error")
		fmt.Println(err)
	}
}
