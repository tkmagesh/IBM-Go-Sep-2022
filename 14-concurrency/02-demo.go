package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //schedule the function execution to the scheduler
	f2()
	time.Sleep(time.Second)
}

func f1() {
	fmt.Println("f1 started")
	//time.Sleep(500 * time.Millisecond)
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
