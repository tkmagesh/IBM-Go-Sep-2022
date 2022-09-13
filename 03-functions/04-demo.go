/* functions as arguments to other functions */
package main

import "fmt"

func main() {
	exec(f1)
	exec(f2)
	exec(func() {
		fmt.Println("anon function invoked")
	})

	/*
		logAdd(100, 200)
		logMultiply(100, 200)
	*/

	logOperation(100, 200, add)
	logOperation(100, 200, multiply)

}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func exec(fn func()) {
	fn()
}

func logOperation(x, y int, oper func(int, int)) {
	fmt.Println("======== operation started =========")
	oper(x, y)
	fmt.Println("======== operation completed =========")
}

func logAdd(x, y int) {
	fmt.Println("======== operation started =========")
	add(x, y)
	fmt.Println("======== operation completed =========")
}

func logMultiply(x, y int) {
	fmt.Println("======== operation started =========")
	multiply(x, y)
	fmt.Println("======== operation completed =========")
}

func add(x, y int) {
	fmt.Printf("Result of adding %d and %d is %d\n", x, y, x+y)
}

func multiply(x, y int) {
	fmt.Printf("Result of multiply %d and %d is %d\n", x, y, x*y)
}
