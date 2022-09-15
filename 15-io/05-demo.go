package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileContents, err := os.ReadFile("data1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(fileContents))
}
