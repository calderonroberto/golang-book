package math //package names generally match the folders they fall in.
// this is a nice convention as it keeps things clean

// To install this package run "go install" in this folder

// public function when importing starts with Capitals

// you can easily generate documentation with:
//godoc github.com/calderonroberto/golang-book/chapter11-packages/math Average

// even better.. you can generate web form docs with:
//godoc -http:=":6060"
// and visiting this url:
// http://localhost:6060/pkg


// Finds the average of a series of numbers (string in doc)
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//private function
func pi() {
	//return float64(3.1515)
}
