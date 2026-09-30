package main

import "fmt"

func oddEven(a int) {
	if a%2 == 0 {
		fmt.Printf("4.  %d is an even number.\n", a)
		return
	}
	fmt.Printf("4.  %d is a odd number.\n", a)
}
