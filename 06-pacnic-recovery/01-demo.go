package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("something went wrong")
			fmt.Println(err)
			return
		}
		fmt.Println("All good")
	}()
	fmt.Println(divide(100, 7))
}

func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	fmt.Println("after calculation")
	return
}
