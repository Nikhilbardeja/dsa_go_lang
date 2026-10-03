package main

import "fmt"

func lcm(num1, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return 0
	}

	fmt.Printf("17. LCM = %d\n", (num1*num2)/gcd(num1, num2))
	return (num1 * num2) / gcd(num1, num2)
}
