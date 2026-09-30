package main

func swapTwoNum(a *int, b *int) {
	var temp int = *b

	*b = *a

	*a = temp

}
