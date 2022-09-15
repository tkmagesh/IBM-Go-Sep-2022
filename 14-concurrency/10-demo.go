package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, wg, ch) //share memory by communicating
	result := <-ch
	wg.Wait()
	fmt.Println("result = ", result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	result := x + y
	ch <- result
}
