package main

import "fmt"

func primeN(num int) {
	fmt.Print("14. Prime numbers: ")
	for i := 0; i <= num; i++ {
		if prime(i) {
			fmt.Printf("%d, ", i)
		}
	}
	fmt.Println()
}
