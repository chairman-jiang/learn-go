package main

import (
	"errors"
	"fmt"
)

// 模拟try catch
func Try(fn func(), handler func(interface{}))  {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fn()
}

var errorMessage = errors.New("second param is required")

func division(x, y int) (int, error) {
	if y == 0 {
		return 0, errorMessage
	}
	return x / y, nil
}

func main() {
	defer func() {
		fmt.Println(recover())
	}()
	switch val, err := division(10, 1); err {
	case nil:
		fmt.Println(val)
	case errorMessage:
		panic(err)
	}
}
