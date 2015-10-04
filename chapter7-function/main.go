package main

import "fmt"

// A global variable 
var msg string =  "Calculating"

// A function called average taking a slice and returning a float64
// The parameters and the return type are the  "function signature"
// func NAME(PARAMETER) RETURNTYPE
func average(xs []float64) float64{
	//panic("Not implemented") //throws an error
	// functions don't have access to calling function vars:
	// fmt.Println(x) //won't work
	// but they have access to global vars
	fmt.Println(msg)
	total := 0.0
	for _,v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

/*
*Functions
*/

// Functions can name the return type:
func f2() (r int) {
	r=1
	return
}

/*
* Returning multiple values
*/

// A function can also return multiple values
// this is often used to return error or success
// like in: x, ok := f()
func f () (int, int) {
	return 5,6
}

/*
* Variadic Functions
*/

// Variadic functions allow us to take zero or more parameters
// by using '...', in this case we take zero or more ints:
func add(args ...int) int{
	total := 0
	for _,v := range args{
		total += v
	}
	return total
}

/*
* Closure
*/

// A function variable with the non-local variables it references
// is known as closure. In main, increment and the variable x form
// the closure. Here is a more interesting example, a function that
// returns another function that when called generates a sequence

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

/*
* Recursion
*/

// A function is able to call itself:
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// In golang it is possible to create functions inside of functions


func main(){
	x := 5
	fmt.Println(x)
	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))

	fmt.Println(f2()) // returning a named type

	x,y := f() //multiple types
	fmt.Println(x,y)

	fmt.Println(add(1,2,3)) // variadic functions
	// we can pass a slice of ints using ...
	xss := []int{1,2,3}
	fmt.Println(add(xss...))

	// you can define function variables that have access
	// to the local variables (same scope) in this function
	increment := func() int {
		x++; return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	// using the closure example above
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
	// unlike normal local variables the local i variable
	// defined in makeEvenGenerator persists between calls

	fmt.Println(factorial(22))

}
