package main

import "fmt"

func main() {
	ch := make(chan int)
	go fn(ch)
	/*
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	*/
	/*
		for i := 0; i < 5; i++ {
			fmt.Println(<-ch)
		}
	*/

	/*
		for {
			if no, ok := <-ch; ok {
				fmt.Println(no)
			} else {
				break
			}
		}
	*/
	for no := range ch {
		fmt.Println(no)
	}
}

func fn(ch chan int) {
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40
	ch <- 50
	close(ch)
}
