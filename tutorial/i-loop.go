package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}
	fmt.Println("Loop completed.")

	// break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Iteration:", i)
	}
	fmt.Println("Loop break completed.")

	// continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Iteration:", i)
	}
	fmt.Println("Loop continue completed.")

	// array
	var fruits = [5]string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Fruit:", fruits[i])
	}
	fmt.Println("Array loop completed.")
	// slice
	var vegetables = []string{"Carrot", "Broccoli", "Spinach"}
	for i := 0; i < len(vegetables); i++ {
		fmt.Println("Vegetable:", vegetables[i])
	}
	fmt.Println("Slice loop completed.")
	// range
	for i, fruit := range fruits {
		fmt.Println("Fruit:", fruit, "at index", i)
	}
	fmt.Println("Range loop completed.")

	// map
	var studentGrades = map[string]int{"John": 85, "Doe": 90, "Jane": 95}
	for name, grade := range studentGrades {
		fmt.Println("Student:", name, "Grade:", grade)
	}
	fmt.Println("Map loop completed.")

	// while
	var count int = 0
	for count < 5 {
		fmt.Println("Count:", count)
		count++
	}
	fmt.Println("While loop completed.")
	// do while
	count = 0
	for {
		fmt.Println("Count:", count)
		count++
		if count >= 5 {
			break
		}
	}
	fmt.Println("Do while loop completed.")



	// infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// 	break
	// }
	// fmt.Println("Infinite loop completed.")
}