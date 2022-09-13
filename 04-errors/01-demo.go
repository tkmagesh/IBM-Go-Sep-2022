/* errors */

package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be 0")

func main() {
	quotient, remainder, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Do not attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, errors.New("divisor cannot be 0")
	}
	return x / y, x % y, nil
}
*/

/*
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("divisor cannot be 0")
		return
	}
	quotient, remainder = x/y, x%y
	return
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient, remainder = x/y, x%y
	return
}
