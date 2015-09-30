package main

import "fmt"

/*
* This example applies what you have learned until this chapter
* including declaring typed variables, using operators, and
* some new bits like string interpolation and input parsing
*/

func main() {

	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

}
