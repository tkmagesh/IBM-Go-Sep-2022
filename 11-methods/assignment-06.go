/*
	Create a Product struct with the following:
		Fields
			Id
			Name
			Cost
			Category
		Methods
			Format()
				return a formatted string representation of the product
					ex: "Id=100, Name=Grapes, Cost=10, Category=Food"
			ApplyDiscount()
				if the cost = 10
				after applying 10% discount
				cost should be 9

*/

package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%q, Cost=%v, Category=%q", p.Id, p.Name, p.Cost, p.Category)
}

func (p *Product) ApplyDiscount(discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry=%q", pp.Product.Format(), pp.Expiry)
}

func NewPerishableProduct(
	id int,
	name string,
	cost float32,
	category string,
	expiry string,
) PerishableProduct {
	return PerishableProduct{
		Product: Product{id, name, cost, category},
		Expiry:  expiry,
	}
}

func main() {
	/*
		grapes := Product{100, "Grapes", 50, "Food"}
		fmt.Println(grapes.Format())

		fmt.Println("After applying 10% discount")
		grapes.ApplyDiscount(10)
		fmt.Println(grapes.Format())
	*/

	grapes := NewPerishableProduct(100, "Grapes", 50, "Food", "2 Days")
	fmt.Println(grapes.Format())

	fmt.Println("After applying 10% discount")
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())
}
