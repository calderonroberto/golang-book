package main

import "fmt"

func zero(x int){
	x = 0
}

func zeroPtr (xPtr *int) { // *int specifies argument must be pointer
	*xPtr = 0
}

func one (xPtr *int) {
	*xPtr = 1
}

func main(){
	x := 5
	zero(x)
	// Calling a function results in the argument being copied to
	// the function. This is good most of the time.
	fmt.Println(x) // x is still 5

	// If we wanted to modify the original function here
	// we can use what is known as a pointer:
	zeroPtr(&x) // &x is the address of the variable x
	fmt.Println(x) // this time x is 0


	// In go a pointer is represented using * followed type (*int)
	// * is used to dereference pointer variables, *xPtr = 0 says
	// "store the int 0 in the memory location xPtr refers to"
	// the & operator is used to find the address of a variable &x
	// so &x in main and xPtr in zero refer to the same memory location


	// Another way to get a pointer is to use new, New takes a type
	// and allocates enough memory to fit a value, returning a pointer
	// Golang is garbage collected, which means when nothing referst
	// to a pointer memory is cleaned up automatically.
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1


}
