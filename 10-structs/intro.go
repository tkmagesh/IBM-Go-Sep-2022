package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
}

func main() {
	//var emp Employee
	//emp := Employee{}
	//emp := Employee{100, "Magesh", "Bengaluru"}
	emp := Employee{Id: 100, Name: "Magesh", City: "Bengaluru"}
	//emp := Employee{Id: 100, Name: "Magesh"}
	//emp := new(Employee)
	fmt.Printf("%#v\n", emp)

	//Accessing the fields
	fmt.Println(emp.City)

	emp2 := &Employee{200, "Suresh", "Pune"}
	fmt.Printf("%#v\n", emp2)
	//fmt.Println((*emp2).Name)
	fmt.Println(emp2.Name)

	/*
		var no int = 10
		noPtr := &no
		fmt.Println(*noPtr)
	*/
}
