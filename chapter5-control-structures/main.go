package main

import "fmt"

func main() {

	/*
	* For Loops
	 */
	// The for statement allows us to repeat a list of statements
	// Notice the lack of parentheses in for loops expressions
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++ //remember to increase i by one
	}

	// We can make a for loop shorter with this pattern:
	// run before ; evaluate each time ; at end run this
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	/*
	* If evaluations
	 */
	// If allows you to evaluate values in a loop
	// let's find even or odd numbers
	for i := 0; i <= 10; i++ {
		if i%2 == 0 { //if the reminder of i/2 equals 0
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	/*
	* Switch
	 */
	// Switch allows you to make many if evaluations more easy
	for i := 10; i >= 0; i-- { //notice how we go from 10 to 0
		switch i {
		//the first one to succeed will be chosen:
		case 0:
			fmt.Println("Zero")
		case 5:
			fmt.Println("Five")
		case 10:
			fmt.Println("Ten")
		default:
			fmt.Println("Unknown Number") //will default
		}
	}

}
