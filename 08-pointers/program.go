package main

import "fmt"

func main() {
	var no int = 10
	//noPtr := &no
	var noPtr *int = &no

	//derefencing - accessing the value using the pointer
	var x = *noPtr
	fmt.Println(x)

	// x == *(&x)

	fmt.Println("Before incrementing, no =", no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)

	n1, n2 := 10, 20
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(z *int) {
	*z += 1
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
