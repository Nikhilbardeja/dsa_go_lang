package main

import "fmt"

func checkPalindromeInt(num int) {
	if num == reverseNum(num) {
		fmt.Printf("7.  Number %d is a palindrome.\n", num)
		return
	}
	fmt.Printf("7.  Number %d is not a palindrome.\n", num)
}
