package main

func factorialRec(num int) int {
	if num == 1 || num == 0 {
		return 1
	}

	return num * factorialRec(num-1)
}

func factorialLoop(num int) int {

	if num == 0 || num == 1 {
		return 1
	}

	var res int = 1

	for j := num; j > 0; j-- {
		res *= j
	}
	return res
}
