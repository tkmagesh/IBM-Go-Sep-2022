package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func (e Employee) WhoAmI() {
	fmt.Println("I am an employee - ", e.Name)
}

func main() {
	emp := Employee{Id: 100, Name: "Magesh", City: "Bengaluru"}
	emp.WhoAmI()
}
