package main

import "fmt"

func main() {
    fmt.Print("Input your number : ")
	var number int
	fmt.Scan(&number)
	if number > 0 {
		fmt.Println("Your number is positive")
	}
	if number < 0 {
		fmt.Println("Your number is negative")
	}
	if number == 0 {
		fmt.Println("Your number is zero")
	}
	if number%2 == 0 {
		fmt.Println("Your number is even")
	} else {
		fmt.Println("Your number is odd")
	}
}