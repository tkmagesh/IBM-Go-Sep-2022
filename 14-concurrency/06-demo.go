package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"math/rand"
)

func main() {
	fmt.Println("Hit ENTER to start")
	fmt.Scanln()
	wg := &sync.WaitGroup{}
	var goRoutineCount = flag.Int("count", 1, "Number of goroutines to be initiated")
	flag.Parse()
	rand.Seed(7)
	for i := 1; i <= *goRoutineCount; i++ {
		wg.Add(1)
		go fn(i, time.Duration(rand.Intn(10))*time.Second, wg)
	}
	wg.Wait()
	fmt.Println("Done")
}

func fn(id int, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn [ id=%d ] started\n", id)
	time.Sleep(delay)
	fmt.Printf("fn [ id=%d ] completed\n", id)
}
