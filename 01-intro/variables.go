package main

import "fmt"

//package level variable
//var msg string = "Hello World! [pkg]"
//var msg = "Hello World! [pkg]"
//msg := "Hello World! [pkg]" //=> invalid syntax at package level

var x = 100

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	//type inference
	/*
		var msg = "Hello World!"
	*/

	msg := "Hello World!" // := syntax is valid ONLY in function scope and NOT in package scope
	fmt.Println(msg)

	//var x = 100

	//multiple variables
	/*
		var x int
		var y int
		var str string
		var result int
		str = "Sum of 100 and 200 is "
		x = 100
		y = 200
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of 100 and 200 is "
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		str = "Sum of 100 and 200 is "
		x = 100
		y = 200
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		str = "Sum of 100 and 200 is "
		x = 100
		y = 200
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			result int    = x + y
			str    string = "Sum of 100 and 200 is "
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of 100 and 200 is "
			result    = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		x, y, str := 100, 200, "Sum of 100 and 200 is "
		result := x + y
		fmt.Println(str, result)
	*/

	x, y := 100, 200
	result := x + y
	str := fmt.Sprintf("Sum of %d and %d is %d\n", x, y, result)
	fmt.Println(str)

	//constants
	const pi float32 = 3.14
	//pi = 2

	//iota
	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 5
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 5
			green
			_
			blue
		)
	*/

	const (
		red = iota * 2
		green
		_
		blue
	)
	fmt.Printf("Red = %d, Green = %d and Blue = %d\n", red, green, blue)

	//bitwise
	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
}
