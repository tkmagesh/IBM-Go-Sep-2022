package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	City string
	Org  *Organization
}

type Organization struct {
	Name string
	City string
}

func main() {
	ibm := &Organization{
		Name: "IBM",
		City: "Bengaluru",
	}
	emp := Employee{
		Id:   100,
		Name: "Magesh",
		City: "Bengaluru",
		Org:  ibm,
	}

	emp2 := Employee{
		Id:   101,
		Name: "Suresh",
		City: "Bengaluru",
		Org:  ibm,
	}

	fmt.Println(emp)
	fmt.Println(emp2)

	fmt.Println("Modifying the city of the organization to Pune")
	ibm.City = "Calcutta"
	//emp.Org.City = "Pune"
	fmt.Println(emp.Org.City)
	fmt.Println(emp2.Org.City)
}
