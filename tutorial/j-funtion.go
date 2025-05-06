package main

import "fmt"

func main() {
	// Function
	showMessage("Hello World")

	average(50, 100)
	avg := average(50, 100)
	fmt.Println("Average:", avg)

	result := total(1, 2)
	fmt.Println("Total:", result)
	
	summationResult, status, oddeven := summation(4, -10)
	fmt.Println("Summation Result:", summationResult, "Status:", status, "Odd/Even:", oddeven)

	// variadic function
	sumResult := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of variadic function:", sumResult)

}
func showMessage(message string) {
	fmt.Println(message)
}

func average(a, b int) float64 {
	return float64(a+b) / 2
}

func total(a int, b int) int {
	return a + b
}

func summation(a, b int) (int, string, string) {
	total := a + b
	status := ""
	oddeven := ""
	if total > 0 {
		status = "positive"
	} else if total < 0 {
		status = "negative"
	} else {
		status = "zero"
	}

	if total%2 == 0 {
		oddeven = "even"
	} else {
		oddeven = "odd"
	}
	
	return total, status , oddeven

}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
