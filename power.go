package main

func power(num, pow int) int {
	if pow == 0 {
		return 1
	}

	return num * power(num, pow-1)
}
