/*
write a goroutine "generatePrimse"  to generate prime number from given start to end
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	end := 100
	primeNoCh := generatePrimes(3, end)
	/* for {
		if primeNo, ok := <-primeNoCh; !ok {
			break
		} else {
			fmt.Println(primeNo)
		}
	} */
	for primeNo := range primeNoCh {
		fmt.Println(primeNo)
	}
	fmt.Println("Done")
}

func generatePrimes(start, end int) chan int {
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				ch <- no
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
