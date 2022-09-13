/*
Write a program that will print all the prime numbers between 3 to 100
*/
package main

import "fmt"

func main() {
	//get the prime numbers from the function
	primeNos := generatePrimes(3, 100)
	//print the prime numbers
	fmt.Println(primeNos)
}

func generatePrimes(start, end int) []int {
	primeNos := make([]int, 0)
LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		primeNos = append(primeNos, no)
	}
	return primeNos
}
