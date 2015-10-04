package main

import "fmt"

func main() {

	/*
	* Arrays
	 */

	// Arrays are numbered sequences of elements,
	//
	var x [5]int // This is how we declare arrays in golang
	x[4] = 100   // This is how we update an element of the array
	// Remember that array addresses begin with 0, e.g. [0 1 2]
	fmt.Println(x)    // Print the whole array
	fmt.Println(x[4]) // Print an element

	// Use an array to calculate the average of test scores
	// A shortcut to create arrays in golang is:
	y := [5]float64{
		98,
		93,
		77,
		82,
		//83, end comma is required, i.e easier to comment out
	}
	var total float64 = 0
	for i := 0; i < len(x); i++ { //len() returns the array length
		total += y[i]
	}
	// len() returns an integer, so we must cast it to float64
	fmt.Println(total / float64(len(x)))

	// We can use a special form of the for loop using range
	var secondtotal float64 = 0 //reset the total to 0
	// '_' tells the compiler we don't need this variable
	for _, value := range y {
		secondtotal += value
	}
	fmt.Println(secondtotal / float64(len(x)))

	/*
	* Slices
	 */

	// Slices are segments of an array, the main difference is
	// their length is allowed to change:
	// var x[]float64
	// x := make([]float64,5) // len(x)=5
	// x := make([]float64,5,10) // len(x)=5, cap(x)=5

	// given: arr := [5]float64{1,2,3,4,5}:

	// x := arr[1:4] // [2,3,4]
	// arr[0:5] == arr[0:len(arr)]
	// arr[:5] == arr[0:5]
	// arr[:] == arr[0:len(arr)]

	// Golang includes two functions: append and copy
	// Append will create a new slice and append to it
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	// Copy slice4 into slice3, since slice4 has only two
	// elements only the first two elements are copied
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	/*
	* Maps
	 */

	// Maps are an unordered collection of key-value pairs
	// they are declared like this:
	// var x map[keyType]valueType, or var x map[string]int
	// Here is how you declare and access a map key
	//var a map[string]int //will cause a runtime error
	a := make(map[string]int)
	a["key"] = 10
	fmt.Println(a)
	fmt.Println(a["key"])

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	// A lookup returns two values, the first one is the lookup
	// the second one if the lookup was successful
	name, ok := elements["Un"] //
	fmt.Println(name, ok)
	if name, ok := elements["Ne"]; ok { //use ok as evaluation
		fmt.Println(name, ok)
	}

	// Let's make a two dimensional map. The first map is
	// used as a lookup table, the inner map to store information
	elementsMap := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}
	if el, ok := elementsMap["H"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}
