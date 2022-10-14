package main

import (
	"errors"
	"fmt"
)

func IPanicEasily() {
	panic("I panic easily")
}

func MyAwesomeFunction() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from a small panic")
			err = errors.New("i panic easily panicked")
		}
		fmt.Println("2")
	}()
	IPanicEasily()
	return nil
}

func main() {
	fmt.Println("Panic! In the Go App")
	if err := MyAwesomeFunction(); err != nil {
		fmt.Println("my awesome function returned an error")
		fmt.Println(err.Error())
	}
	fmt.Println("finished main function successfully")
}
