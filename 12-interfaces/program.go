package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

type ShapeWithArea interface {
	Area() float32
}

func printArea(sa ShapeWithArea) {
	fmt.Println("Area = ", sa.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func printPerimeter(sp ShapeWithPerimeter) {
	fmt.Println("Perimeter = ", sp.Perimeter())
}

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func printShape(s Shape) {
	printArea(s)
	printPerimeter(s)
}

func main() {
	c := Circle{12}
	//fmt.Println("Area = ", c.Area())
	/*
		printArea(c)
		printPerimeter(c)
	*/
	printShape(c)

	r := Rectangle{10, 12}
	//fmt.Println("Area = ", r.Area())
	/*
		printArea(r)
		printPerimeter(r)
	*/
	printShape(r)
}
