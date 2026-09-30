package main

import "fmt"

func sumOfDig(num int) {
	var sum int = 0
	for num != 0 {
		sum += num % 10
		num /= 10
	}

	fmt.Printf("10. Sum of digits is %d.\n", sum)

}
