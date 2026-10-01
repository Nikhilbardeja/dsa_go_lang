package main

import "fmt"

func countFactors(num int) {
	var count, factor, max int = 2, 2, num / 2
	for factor <= max {
		if num%factor == 0 {
			count++
		}
		factor++

	}

	fmt.Printf("15. Factors = %d\n", count)
}
