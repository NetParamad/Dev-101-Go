package main

import "fmt"

func main() {

	var daysOfWeek [7]string = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	var count int = len(daysOfWeek)
	fmt.Println("Days of the week:", daysOfWeek)
	fmt.Println("Number of days in the week:", count)

	student := [...] string{"John", "Doe", "Jane", "Smith", "Emily"}
	fmt.Println("Student names:", student)

    var arr [5]int
    fmt.Println("Enter 5 integers:")
    for i := 0; i < 5; i++ {
		fmt.Print("Enter integer ", i+1, ": ")
        fmt.Scan(&arr[i])
    }
    fmt.Println("You entered:", arr)
    fmt.Println("The sum of the entered integers is:", sum(arr))
}

func sum(arr [5]int) int {
    total := 0
    for _, value := range arr {
        total += value
    }
    return total
}