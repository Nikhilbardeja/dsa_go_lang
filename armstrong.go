package main

import "fmt"

func armstrong(num int) int {
	var result, count int = 0, countDig(num)

	for i := 0; i < count; i++ {
		result += power(num%10, count)
		num /= 10

	}

	fmt.Printf("18. Armstrong number = %d\n", result)
	return result
}
