package main

import "fmt"

func main() {

	studentGrades := map[string]int{
		"John": 85,
		"Doe":  90,
		"Jane": 95,
		"Smith": 80,
		"Emily": 88,
	}
	fmt.Println("Student Grades:", studentGrades)

	fmt.Println("John's grade:", studentGrades["John"])

	studentGrades["Alice"] = 92
	fmt.Println("Updated Student Grades:", studentGrades)
	delete(studentGrades, "Doe")
	fmt.Println("After deletion:", studentGrades)
	fmt.Println("Total number of students:", len(studentGrades))
	fmt.Println("All student names:")
	for name := range studentGrades {
		fmt.Println(name)
	}
	fmt.Println("All student grades:")
	for _, grade := range studentGrades {
		fmt.Println(grade)
	}

}