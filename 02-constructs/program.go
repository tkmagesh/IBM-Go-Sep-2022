package main

import "fmt"

func main() {
	//if
	fmt.Print("\nif else\n")
	/*
		x := 21
		if x%2 == 0 {
			fmt.Printf("%d is an even number\n", x)
		} else {
			fmt.Printf("%d is an odd number\n", x)
		}
	*/

	if x := 21; x%2 == 0 { //scope of the variable 'x' is limited to the 'if' block
		fmt.Printf("%d is an even number\n", x)
	} else {
		fmt.Printf("%d is an odd number\n", x)
	}

	//switch case
	fmt.Print("\nSwitch case\n")

	/*
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very Good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/

	score := 6
	/*
		switch score {
		case 0:
			fmt.Println("Terrible")
		case 1:
			fmt.Println("Terrible")
		case 2:
			fmt.Println("Terrible")
		case 3:
			fmt.Println("Terrible")
		case 4:
			fmt.Println("Not Bad")
		case 5:
			fmt.Println("Not Bad")
		case 6:
			fmt.Println("Not Bad")
		case 7:
			fmt.Println("Not Bad")
		case 8:
			fmt.Println("Very Good")
		case 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid score")

		}
	*/
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	/*
		x := 16
		switch {
		case x%2 == 0:
			fmt.Printf("%d is an even number\n", x)
		case x%2 != 0:
			fmt.Printf("%d is an odd number\n", x)
		}
	*/

	switch x := 16; {
	case x%2 == 0:
		fmt.Printf("%d is an even number\n", x)
	case x%2 != 0:
		fmt.Printf("%d is an odd number\n", x)
	}

	fmt.Println("Switch case with fallthrough")
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	case 6:
		fmt.Println("is <= 6")
		//fallthrough
	case 7:
		fmt.Println("is <= 7")
		fallthrough
	case 8:
		fmt.Println("is <= 8")
	}

	plan := "Super"
	switch plan {
	case "Premium":
		fmt.Println("All Premium features")
		fallthrough
	case "Super":
		fmt.Println("All Super features")
		fallthrough
	case "Pro":
		fmt.Println("All Pro features")
		fallthrough
	case "Free":
		fmt.Println("All Free features")
	default:
		fmt.Println("Invalid plan")
	}
	fmt.Println()
}
