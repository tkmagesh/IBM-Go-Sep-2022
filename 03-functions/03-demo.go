package main

import "fmt"

func main() {

	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var getGreetMsg func(string) string
	getGreetMsg = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
	}
	print(getGreetMsg("Suresh"))

	var add ?
	add = func(x, y int) (result int) {
		result = x + y
		return
	}
	fmt.Println(add(100,200))

	var divide ?
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	fmt.Println(divide(100,7))
}
