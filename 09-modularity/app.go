package main

import (
	"fmt"
	"modularity-demo/calculator"
	"modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	//fmt.Println("Module executed")
	color.Red("Module executed")

	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("OpCount =", calculator.OpCount())
	fmt.Println(utils.IsEven(11))
}
