package main

import "fmt"

func main() {
	// Slices
	var fruits []string = []string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Println("Fruits:", fruits)
	

	fruits = append(fruits, "Elderberry")
	fmt.Println("Updated Fruits:", fruits)
	fmt.Println("Number of fruits:", len(fruits))
	fmt.Println("First fruit:", fruits[0])
	fmt.Println("Last fruit:", fruits[len(fruits)-1])
	fmt.Println("Slice of fruits:", fruits[1:3]) // Slice from index 1 to 3 (exclusive)
	fmt.Println("Slice of fruits:", fruits[:3])
	fmt.Println("Slice of fruits:", fruits[1:])
	fmt.Println("Slice of fruits:", fruits[:])
}