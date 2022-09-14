package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

//fmt.Stringer interface implementation
func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func main() {
	pen := Product{100, "Pen", 10, 100, "Stationary"}
	fmt.Println(pen)
}
