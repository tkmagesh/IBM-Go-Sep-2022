package main

import (
	"fmt"
	"modularity-demo/calculator"
)

func main() {
	fmt.Println("Module executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("OpCount =", calculator.OpCount())
}
