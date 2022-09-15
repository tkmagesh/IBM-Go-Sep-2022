package main

import (
	"fmt"
)

func main() {
	go f1() //schedule the function execution to the scheduler
	f2()
	//time.Sleep(time.Second)
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
