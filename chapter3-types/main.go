package main

/*
Types
*/

import "fmt"

func main() {

	// Integers; will print "1".
	fmt.Println("1 + 1 = ", 1+1)

	// Floats; will also print "1".
	fmt.Println("1 + 1 = ", 1.0+1.0)

	// String manipulation
	// Double quotes allow escape characters like \n
	// Single back ticks allow newlines
	fmt.Println(`Hello world,
		What up!`) //will print new line and tabs
	fmt.Println(len("Hello World")) //get the length
	fmt.Println("Hello World"[1])   //access a string byte (will print 101)
	fmt.Println("Hello " + "World") //concatenate

	// Booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
