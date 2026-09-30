package main

import "fmt"

func checkPalindromeInt(num int) {
	if num == reverseNum(num) {
		fmt.Printf("7.  Number %d is a palindrome.\n", num)
		return
	}
	fmt.Printf("7.  Number %d is not a palindrome.\n", num)
}

func checkPalindromeStr(num string) {
	var i, length int = len(num) / 2, len(num)

	for j := 0; j < i; j++ {
		if num[j] != num[length-1-j] {
			fmt.Printf("7.  Number %s is not a palindrome.\n", num)
			return
		}
	}
	fmt.Printf("7.  Number %s is a palindrome.\n", num)

}
