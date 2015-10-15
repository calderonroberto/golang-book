package math

import "testing"

// Tests are identified by starting a function with the work Test
// and taking one argument of type *testing.T, convention is to 
// use the name of the funciton tested (i.e. TestAverage)

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1,2})
	if v != 1.5{
		t.Error("Expected 1.5, got ", v)
	}
}
