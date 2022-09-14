package main

import (
	"errors"
	"fmt"
)

func main() {

	quotient, remainder, err := divideClient(100, 0)
	if err != nil {
		fmt.Println("Divide operation failed")
		fmt.Println("Perform an alternate operation")
	} else {
		fmt.Println(quotient, remainder)
	}
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

//from a 3rd party library
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(errors.New("divisor cannot be zero"))
	}
	quotient, remainder = x/y, x%y
	fmt.Println("after calculation")
	return
}
