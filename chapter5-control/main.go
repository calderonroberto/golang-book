package main

import "fmt"

func main() {
 	i := 1

	//notice the lack of parentheses in for loops expressions
 	for i <= 10 {
		fmt.Println(i)
		i = i + 1
 	}
}

