package main

import (
	"flag"
	"fmt"
	"math/rand"
)

// you can invoke go run flags.go -max=100

func main() {
	//Define flags, with default value
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))

	// Any additional non-flags arguments can be retrieved with
	// flag.Args(), which returns []string

}
