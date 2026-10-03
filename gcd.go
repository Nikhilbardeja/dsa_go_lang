package main

import "fmt"

func gcd(num1 int, num2 int) int {
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	fmt.Printf("16. GCD = %d\n", num1)
	return num1
}
