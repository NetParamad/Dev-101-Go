package main

import "fmt"
func main() {

	fmt.Println("Input your name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello", name)
	fmt.Println("Input your age:")
	var age int
	fmt.Scanln(&age)
	fmt.Println("You are", age, "years old")
	fmt.Println("Input your height:")
	var height float64
	fmt.Scanln(&height)
	fmt.Println("Your height is", height, "cm")
}