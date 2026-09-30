package main

import "fmt"

func largestOfThree(a int, b int, c int) int {
	var largest int = a

	if largest < b {
		largest = b
	}

	if largest < c {
		largest = c
	}

	fmt.Printf("5.  %d is largest among %d, %d and %d\n", largest, a, b, c)

	return c
}
