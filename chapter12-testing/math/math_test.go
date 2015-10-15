package math

import "testing"

// Tests are identified by starting a function with the work Test
// and taking one argument of type *testing.T, convention is to 
// use the name of the funciton tested (i.e. TestAverage)


// Struct to represent input and output for the function
type testpair struct {
	values[]float64
	average float64
}

// A list of inputs and expected outputs
var tests = []testpair{
	{ []float64{1,2}, 1.5 },
	{ []float64{1,1,1,1,1,1}, 1 },
	{ []float64{-1,1}, 0 },
}


// Test that for each in the list of inputs/outputs.
func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v!= pair.average{
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
