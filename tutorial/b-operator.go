package main

import "fmt"



func main() {
// operator
// +, -, *, /, %, ++, --, +=, -=, *=, /=, %=
// &&, ||, !, ==, !=, >, <, >=, <=
// bitwise operator
// &, |, ^, <<, >>, &^
// & = and, | = or, ^ = xor, << = left shift, >> = right shift, &^ = and not

var a int = 10
var b int = 20

fmt.Println("sum of a and b:", a + b)
fmt.Println("difference of a and b:", a - b)
fmt.Println("product of a and b:", a * b)
fmt.Println("quotient of a and b:", a / b)
fmt.Println("remainder of a and b:", a % b)

}