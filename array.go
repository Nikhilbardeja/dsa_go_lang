package main

import (
	"fmt"
	"math"
)

func maxArray(arr []int) int {
	var max int = arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	fmt.Printf("31.  Max = %d\n", max)
	return max
}

func minArray(arr []int) int {
	var min int = arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	fmt.Printf("32.  Min = %d\n", min)
	return min
}

func sumArray(arr []int) int {
	var sum int = 0
	for _, value := range arr {
		sum += value
	}
	fmt.Printf("33.  Sum = %d\n", sum)
	return sum
}

func avgArray(arr []float32) float32 {
	var sum float32 = 0
	for _, value := range arr {
		sum += value
	}
	fmt.Printf("34.  Avg = %f\n", sum/float32(len(arr)))
	return sum / float32(len(arr))
}

func linSearch(arr []int, ele int) int {

	for i, value := range arr {
		if value == ele {
			fmt.Printf("35.  Avg = %d\n", i)
			return i
		}
	}
	return -1
}

func revArray(arr []int) []int {
	var l int = len(arr)
	for i := 0; i < l/2; i++ {
		var temp int = arr[i]
		arr[i] = arr[l-i-1]
		arr[l-i-1] = temp
	}
	fmt.Printf("36.  Rev Array = %d\n", arr)
	return arr
}

func secMaxArray(arr []int) int {
	var l1, l2 int = math.MinInt, math.MinInt
	for _, value := range arr {
		if value > l1 {
			l2 = l1
			l1 = value
		} else if value > l2 && value != l1 {
			l2 = value
		}
	}
	fmt.Printf("37.  Sec Max = %d\n", l2)
	return l2
}

func countOddEvenArray(arr []int) (int, int) {
	var odd, even int = 0, 0
	for _, value := range arr {
		if value%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}
	fmt.Printf("38.  Even nums = %d, Odd nums = %d\n", even, odd)
	return even, odd
}

func removeDupsArray(arr []int) {
	result := []int{0}

	result[0] = arr[0]
	i := 0
	for _, value := range arr {
		if result[i] != value {
			result = append(result, value)
			i++
		}
	}
	fmt.Printf("39.  Remove Duplicates Array = %d\n", result)

}

func rightRotate(arr []int) {
	result := make([]int, len(arr))
	result[0] = arr[len(arr)-1]

	for i := 1; i < len(arr); i++ {
		result[i] = arr[i-1]
	}
	fmt.Printf("40.  Rotate Right Array = %d\n", result)
}

// {4, 0, 1, 2, 3, 5, 4}

// {0, 1, 2, 3, 4, 5}
