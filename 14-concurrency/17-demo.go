/* modify this program to generate prime numbers for 20 seconds */

package main

import (
	"fmt"
	"time"
)

func main() {

	primeNoCh := generatePrimes(3)
	for primeNo := range primeNoCh {
		fmt.Println(primeNo)
	}
	fmt.Println("Done")
}

func generatePrimes(start int) chan int {
	ch := make(chan int)
	timeoutCh := time.After(20 * time.Second)
	//timeoutCh := timeout(20 * time.Second)
	go func() {
		no := start
	LOOP:
		for {
			if !isPrime(no) {
				no++
				select {
				case <-timeoutCh:
					break LOOP
				default:
					continue LOOP
				}
			}
			select {
			case <-timeoutCh:
				break LOOP
			case ch <- no:
				no++
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= no/2; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func timeout(d time.Duration) <-chan time.Time {
	timeOutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeOutCh <- time.Now()
	}()
	return timeOutCh
}
