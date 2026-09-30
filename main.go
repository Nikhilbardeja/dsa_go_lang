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

	leapYear(2024)
	leapYear(1900)
	leapYear(2000)

	reverseNum(123)

	checkPalindromeInt(12321)
	checkPalindromeInt(65126)

	checkPalindromeStr("12321")
	checkPalindromeStr("65321")

	countDig(1234)
	countDig(0)

	sumOfDig(12345)
	sumOfDig(0)

}
