package main

import "fmt"

func countDig(num int) {
	var i int = 0
	for num != 0 {
		num /= 10
		i++
	}

	fmt.Printf("9.  This number has %d digits.\n", i)

}
