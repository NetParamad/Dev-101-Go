package main

import "fmt"

// ตัวแปร global
var name  string = "NetParamad"
const pi float64 = 3.14

func main(){
	// ตัวแปร local
	var age int = 20
	var grade float64 = 3.5

	// ตัวแปร inferred type
	isStudent := true
	
	fmt.Println("Hello, World!")
	fmt.Println("my name is", name , "and I'm", age, "years old" , "and my grade is", grade)
	fmt.Println("Am I a student?", isStudent)
	fmt.Println("Pi is", pi)

	fmt.Printf("my name is %s and I'm %v years old and my grade is %.2f\n", name, age, grade)// %s = string, %d = int, %.2f = float and %v = value
	// data type
	fmt.Printf("my name is %T and I'm %T years old and my grade is %T\n", name, age, grade)// %T = type
	

}