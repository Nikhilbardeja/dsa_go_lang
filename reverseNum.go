package main

import "fmt"

func reverseNum(num int) {
	var result int = 0

	for num != 0 {

		var dig int = num % 10

		result = result*10 + dig

		num /= 10

	}

	fmt.Printf("Reveresed Num = %d\n", result)
}
