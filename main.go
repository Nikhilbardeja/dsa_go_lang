package main

import "fmt"

func main() {
	helloWorld()

	var a, b int = 3, 4
	sumTwoNum(a, b)

	fmt.Printf("3.  Before: a = %d , b= %d \n", a, b)
	swapTwoNum(&a, &b)
	fmt.Printf("3.  After: a = %d , b= %d \n", a, b)

	oddEven(0)
	oddEven(2)
	oddEven(3)

	largestOfThree(1, 2, 3)
	largestOfThree(5, 4, 3)
	largestOfThree(3, 6, 4)

}
