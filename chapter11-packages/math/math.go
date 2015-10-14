package math //package names generally match the folders they fall in.
// this is a nice convention as it keeps things clean

// To install this package run "go install" in this folder

//public function when importing
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//private function
func pi() {
	return float64(3.1515)
}
