package main

import "fmt"

//use an alias "m" to not conflict with std math package
import m "github.com/calderonroberto/golang-book/chapter11-packages/math"

//notice how in my case i need to use github.com/username, as it's part of a github account

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
