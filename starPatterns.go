package main

import "fmt"

func first(num int) {
	fmt.Println("21.  ")
	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func second(num int) {
	fmt.Println("22.  ")

	for i := num; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func third(num int) {
	fmt.Println("23.  ")
	for i := 0; i < num; i++ {
		for j := 1; j <= i+1; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func fourth(num int) {
	var s string = "ABCDEFG"
	fmt.Println("24.  ")
	for i := 0; i < num; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%c", s[j])
		}
		fmt.Println()
	}

}

func five(num int) {
	fmt.Println("25.  ")
	for i := 1; i <= num; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func six(num int) {
	fmt.Println("26.  ")

	for i := 0; i < num; i++ {
		for j := 0; j < num+i-1; j++ {
			if j < num-i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func seven(num int) {
	fmt.Println("27.  ")

	for i := num; i > 0; i-- {
		for j := 0; j < num+i-1; j++ {
			if j < num-i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func eight(num int) {
	fmt.Println("28.  ")
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			if i == 1 || j == 1 || i == num || j == num {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func nine(num int) {
	fmt.Println("29.  ")
	for i := 0; i <= num; i++ {

		for j := 1; j <= num-i; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == num {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func ten(num int) {
	fmt.Println("30.  ")

	for i := 0; i <= num; i++ {
		for j := 0; j <= 2*num; j++ {
			if j <= i || j >= 2*num-i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	for i := num; i >= 0; i-- {
		for j := 0; j <= 2*num; j++ {
			if j <= i || j >= 2*num-i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}
