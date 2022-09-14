package main

import "fmt"

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	s := MyStr("this is a sample string")
	fmt.Println(s.Length())
}
