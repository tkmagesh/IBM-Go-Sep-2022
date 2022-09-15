package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go f1() //schedule the function execution to the scheduler
	wg.Add(1)
	go f2()
	wg.Wait() // blocking until the wg counter becomes 0
}

func f1() {
	//fmt.Println("f1 invoked")
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
	wg.Done()
}
