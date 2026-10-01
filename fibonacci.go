package main

import "fmt"

func fibonacci() {

	fmt.Print("12. Fibonacci 0, 1, ")
	var i, j int = 0, 1

	for j < 100 {
		k := i + j
		fmt.Printf("%d, ", k)

		i = j
		j = k

	}
	fmt.Println()
}
