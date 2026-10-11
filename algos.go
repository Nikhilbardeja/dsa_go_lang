package main

import "fmt"

func binarySearch(arr []int, num int) int {
	return binarySearchHelper(arr, num, 0, len(arr)-1)
}

func binarySearchHelper(arr []int, num int, low int, high int) int {
	if low > high {
		return -1
	}
	var mid int = low + (high-low)/2

	if arr[mid] == num {
		return mid
	} else if arr[mid] > num {
		return binarySearchHelper(arr, num, 0, mid-1)
	} else {
		return binarySearchHelper(arr, num, mid+1, high)
	}
}

func bubbleSort(arr []int) {
	var n int = len(arr) - 1

	for i := 0; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	fmt.Printf("47. Sorted Array: %d\n", arr)
}

func selectionSort(arr []int) {
	var n int = len(arr) - 1

	for i := 1; i < n; i++ {
		min := arr[i]
		minIndex := i

		for j := i + 1; j <= n; j++ {
			if min > arr[j] {
				min = arr[j]
				minIndex = j
			}
		}
		temp := arr[i]
		arr[i] = min
		arr[minIndex] = temp

	}
	fmt.Printf("48. Selection Sorted Array: %d\n", arr)
}

func insertionSort(arr []int) {
	var n int = len(arr)

	for i := 1; i < n; i++ {
		ele := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > ele {
			arr[j+1] = arr[j]
			j--
		}

		// Insert the element into its correct position
		arr[j+1] = ele
	}
	fmt.Printf("49. Insertion Sorted Array: %d\n", arr)
}

func mergeSortedArrays(arr1, arr2 []int) []int {
	n1, n2 := len(arr1), len(arr2)
	result := make([]int, 0, n1+n2)

	i, j := 0, 0

	// Compare elements from both arrays and add the smaller one
	for i < n1 && j < n2 {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	// If there are any elements left in arr1, append them
	for i < n1 {
		result = append(result, arr1[i])
		i++
	}

	// If there are any elements left in arr2, append them
	for j < n2 {
		result = append(result, arr2[j])
		j++
	}
	fmt.Printf("50. Merged Array: %d\n", result)
	return result
}
