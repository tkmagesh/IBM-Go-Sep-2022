/*
	Display the following menu
		1. Add
		2. Subtract
		3. Multiply
		4. Divide
		5. Exit
	Accept the user choice
	If the user choice is 1 - 4
		accept 2 numbers from the user
		perform the corresponding operation
		print the result
		display the menu again

	If the user choice = 5
		print "Thank you"
		exit the application

	If the user choice is NOT 1 - 5
		print "Invalid choice"
		display the menu again
*/

package main

import "fmt"

func main() {
	var userChoice, result int
	for {
		userChoice = getUserChoice()
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		result = processChoice(userChoice)
		fmt.Println("result =", result)
	}
	fmt.Println("Thank you!")
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice:")
	_, err := fmt.Scanln(&userChoice)
	if err != nil {
		fmt.Println(err)
	}
	return userChoice
}

func getOperands() (n1, n2 int) {
	fmt.Println("Enter the numbers : ")
	fmt.Scanln(&n1, &n2)
	return
}

func processChoice(userChoice int) int {
	var result int
	n1, n2 := getOperands()
	switch userChoice {
	case 1:
		result = add(n1, n2)
	case 2:
		result = subtract(n1, n2)
	case 3:
		result = multiply(n1, n2)
	case 4:
		result = divide(n1, n2)
	}
	return result
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
