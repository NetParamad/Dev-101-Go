package main

import "fmt"

func main() {
	// switch case

	var number int
	fmt.Print("Input your number : ")
	fmt.Scan(&number)
	switch {
	case number > 0:
		fmt.Println("Your number is positive")
	case number < 0:
		fmt.Println("Your number is negative")
	case number == 0:
		fmt.Println("Your number is zero")
	default:
		fmt.Println("Your number is not a number")
	}
}