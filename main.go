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

	countDig(1234)
	countDig(0)

	sumOfDig(12345)
	sumOfDig(0)

	fmt.Printf("11. Factorial = %d\n", factorialRec(5))
	fmt.Printf("11. Factorial = %d\n", factorialLoop(5))

	fibonacci()

	prime(10)
	prime(19)

	primeN(10)

	countFactors(16)

	gcd(48, 18)

	lcm(3, 6)

	armstrong(153)

	fmt.Printf("19. Perfect Num(6) = %t\n", perfectNum(6))
	fmt.Printf("19. Perfect Num(8) = %t\n", perfectNum(8))

	fmt.Printf("20. 3^3 = %d\n", power(3, 3))

	first(5)
	second(5)
	third(5)
	fourth(5)
	five(5)
	six(5)
	seven(5)
	eight(5)
	nine(5)
	ten(5)

	maxArray([]int{1, 2, 34, 43, 45, 3, 234, 25, 45, 23, 23})
	minArray([]int{1, 2, 34, 43, 45, 3, 234, 25, 45, 23, 23, -8})
	sumArray([]int{1, 2, -3, 1})
	avgArray([]float32{1, 2, -3, 1})
	linSearch([]int{1, 2, -3, 1}, -3)
	revArray([]int{0, 1, 2, 3, 5, 4})
	secMaxArray([]int{0, 0, 1, 2, 3, 5, 4})
	countOddEvenArray([]int{0, 0, 1, 2, 3, 5, 4})
	removeDupsArray([]int{0, 0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 5, 5})
	rightRotate([]int{1, 2, 3, 4, 5})

	revStr("ABC")
	checkPalindromeStr("12321")
	checkPalindromeStr("65321")

	countVowelConsonant("aiouebcdf")

	freqencyOfChar("aabcdj")

	removeSpaces("a c b d ")
}
