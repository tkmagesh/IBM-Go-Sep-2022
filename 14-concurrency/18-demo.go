/* modify this program to generate prime numbers for 20 seconds */

package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan bool)
	primeNoCh := generatePrimes(3, stopCh)
	fmt.Println("Hit ENTER key to stop!")
	go func() {
		fmt.Scanln()
		stopCh <- true
	}()
	for primeNo := range primeNoCh {
		fmt.Println(primeNo)
	}
	fmt.Println("Done")
}

func generatePrimes(start int, stopCh chan bool) chan int {
	ch := make(chan int)
	//timeoutCh := time.After(20 * time.Second)
	go func() {
		no := start
	LOOP:
		for {
			if !isPrime(no) {
				no++
				select {
				case <-stopCh:
					break LOOP
				default:
					continue LOOP
				}
			}
			select {
			case <-stopCh:
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
