package main

import "fmt"

func main() {
	//var nos []int
	//var nos []int = []int{3, 1, 4, 2, 5}
	//nos := []int{3, 1, 4, 2, 5}
	var nos []int
	nos = make([]int, 5)
	nos[0] = 3
	nos[1] = 1
	nos[2] = 4
	nos[3] = 2
	nos[4] = 5
	fmt.Printf("%#v\n", nos)

	//nos = append(nos, 100)
	//nos = append(nos, 100, 200, 300)
	hundreds := []int{100, 200, 300, 400}
	nos = append(nos, hundreds...)
	fmt.Printf("%#v\n", nos)

	fmt.Println("Iterating a slice")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println()
	fmt.Println("Iterating a slice (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	/*
		slicing
		[lo:hi] => from lo to hi-1
		[lo:] => from lo to end of the slice
		[:hi] => from 0 to hi - 1
	*/

	fmt.Println("nos[3:6] =", nos[3:6])
	fmt.Println("nos[3:] =", nos[3:])
	fmt.Println("nos[:6] =", nos[:6])

}
