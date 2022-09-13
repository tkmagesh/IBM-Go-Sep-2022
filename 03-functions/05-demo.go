package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		fn := getFn()
		fn()
		getFn()()
	*/

	/*
		logAdd := logOperation(add)
		logAdd(100, 200)

		logMultiply := logOperation(multiply)
		logMultiply(100, 200)
	*/

	/*
		profileAdd := profileOperation(add)
		profileAdd(100, 200)

		profileMultiply := profileOperation(multiply)
		profileMultiply(100, 200)
	*/
	/*
		logAdd := logOperation(add)
		profileAndLogAdd := profileOperation(logAdd)
		profileAndLogAdd(100, 200)
	*/
	/*
		profileAndLogAdd := profileOperation(logOperation(add))
		profileAndLogAdd(100, 200)
	*/
	profileOperation(logOperation(add))(100, 200)
	fmt.Println()

	logMultiply := logOperation(multiply)
	profileAndLogMultiply := profileOperation(logMultiply)
	profileAndLogMultiply(100, 200)
}

func add(x, y int) {
	fmt.Printf("Result of adding %d and %d is %d\n", x, y, x+y)
}

func multiply(x, y int) {
	fmt.Printf("Result of multiply %d and %d is %d\n", x, y, x*y)
}

func logOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("======== operation started =========")
		oper(x, y)
		fmt.Println("======== operation completed =========")
	}
}

func profileOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		oper(x, y)
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Println("Time taken = ", elapsed)
	}
}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}
