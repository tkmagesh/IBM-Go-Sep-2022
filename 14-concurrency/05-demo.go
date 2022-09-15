package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		var wg sync.WaitGroup
		wg.Add(1)
		go f1(&wg) //schedule the function execution to the scheduler
		wg.Add(1)
		go f2(&wg)
	*/
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go f1(wg) //schedule the function execution to the scheduler
	wg.Add(1)
	go f2(wg)
	wg.Wait() // blocking until the wg counter becomes 0
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println("f1 invoked")
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	return
	//wg.Done() // decrement the wg counter by 1
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f2 invoked")

}
