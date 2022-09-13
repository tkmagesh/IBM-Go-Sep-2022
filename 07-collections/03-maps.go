package main

import "fmt"

func main() {
	//var productRanks map[string]int

	/*
		productRanks := make(map[string]int)
		productRanks["pen"] = 5
		productRanks["pencil"] = 1
		productRanks["marker"] = 3
	*/

	//productRanks := map[string]int{"pen": 5, "pencil": 1, "marker": 3}

	productRanks := map[string]int{
		"pen":    5,
		"pencil": 1,
		"marker": 3,
	}
	productRanks["notebook"] = 2
	fmt.Println(productRanks)

	fmt.Println()
	fmt.Println("Accessing the value by key")
	fmt.Printf("Rank of %q is %d\n", "pencil", productRanks["pencil"])

	fmt.Println()
	fmt.Println("Iterating a map")
	for product, rank := range productRanks {
		fmt.Printf("productRanks[%s] = %d\n", product, rank)
	}

	fmt.Println()
	fmt.Println("deleting an entry")
	delete(productRanks, "marker")
	fmt.Println("After deleting 'marker'")
	fmt.Println(productRanks)

	fmt.Println()
	fmt.Println("Check if a key exists")
	keyToSearch := "scribble-pad"
	if rank, exists := productRanks[keyToSearch]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToSearch, rank)
	} else {
		fmt.Printf("%q does not exist\n", keyToSearch)
	}
}
