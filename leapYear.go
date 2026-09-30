package main

import "fmt"

func leapYear(year int) {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Printf("6.  Year %d is a leap year\n", year)
		return
	}
	fmt.Printf("6.  Year %d is not a leap year\n", year)
}
