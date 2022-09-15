package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*
	Write the apis for the following
		Sort => Sort the products collection by any attribute
			IMPORTANT : MUST Use sort.Sort() function
            use cases:
                1. sort the products collection by cost
                2. sort the products collection by name
                3. sort the products collection by units
                4. sort the products collection by cost in descending order
                5. sort the products collection by name in descending order
                6. sort the products collection by units in descending order
				etc

*/

type Products []Product

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result = fmt.Sprintf("%s%v\n", result, product)
	}
	return result
}

/* sort.Interface interface implementation */
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

/* sorting by name */
type byName struct {
	Products
}

func (products byName) Less(i, j int) bool {
	return products.Products[i].Name < products.Products[j].Name
}

/* sorting by cost */
type byCost struct {
	Products
}

func (products byCost) Less(i, j int) bool {
	return products.Products[i].Cost < products.Products[j].Cost
}

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(byName{products})
	case "Cost":
		sort.Sort(byCost{products})
	case "Units":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Units < products[j].Units
		})
	}
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Default sort")
	//sort.Sort(products)
	products.Sort("Id")
	fmt.Println(products)

	fmt.Println("sort by name")
	//sort.Sort(byName{products})
	products.Sort("Name")
	fmt.Println(products)

	fmt.Println("sort by cost")
	//sort.Sort(byCost{products})
	products.Sort("Cost")
	fmt.Println(products)

	fmt.Println("sort by units")
	products.Sort("Units")
	fmt.Println(products)

}
