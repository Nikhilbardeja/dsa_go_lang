package main

func prime(num int) bool {
	var i int = num/2 + 1

	for j := 2; j <= i; j++ {

		if num%j == 0 {
			// fmt.Println("13. Number is prime")
			return true
		}
	}

	// fmt.Println("13. Number is not prime")
	return false
}
