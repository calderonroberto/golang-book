package main

import "fmt"

var globalVariable = "Global Variable"


func main() {

	/*
	* Introduction
	*/

	// Declaring a variable with "var" and its type.
	var x string 
	x = "Hello"
	fmt.Println(x)
	x += " there." // Concatenating strings
	var y = x + " Sup!"
	fmt.Println(x)
	fmt.Println(y)

	// Remember that == is a boolean operator
	fmt.Println(x == y)
	var a string = "hello"
	var b string = "hello"
	fmt.Println(a == b)

	// Declaring a variable can be done (shorter) like:
	c := "Hello Martians!"
	fmt.Println(c)

	// Declaring a variable with "var" and no type
	var d = "Hello Neptunians!"
	fmt.Println(d)

     	/*
        * Naming Conventions
        */

	// A convention in golang are "camel case" variable names
	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)

	/*
	* Scope
	*/

	// Variables outside can be accessed
	fmt.Println(globalVariable)

	/*
	* Constants
	*/

	const constantString string = "A Constant String"
	// constantString = "Will cause compile time error"
	fmt.Println(constantString)

	/*
	* Multiple variable definition
	*/

	// A shorthand to define multiple variables
	var (
	  s = 5
	  u = 10
	  p = 15
	)
	fmt.Println(s + u - p)
}

func f() {
	// This will print "undefined: x" because the variable
	// x is not defined within this function (it is in main)
	// fmt.Println(x)
}
