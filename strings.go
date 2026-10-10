package main

import (
	"fmt"
	"unicode"
)

func revStr(str string) string {
	var bytes []byte = []byte(str)
	var l int = len(bytes)
	for i := 0; i < l/2; i++ {
		var temp byte = bytes[i]
		bytes[i] = str[l-i-1]
		bytes[l-i-1] = temp
	}
	str = string(bytes)
	fmt.Printf("41.  Rev String = %s\n", str)
	return str
}

func checkPalindromeStr(num string) {
	var i, length int = len(num) / 2, len(num)

	for j := 0; j < i; j++ {
		if num[j] != num[length-1-j] {
			fmt.Printf("42.  Number %s is not a palindrome.\n", num)
			return
		}
	}
	fmt.Printf("42.  Number %s is a palindrome.\n", num)

}

func countVowelConsonant(str string) {
	var vowels, consonants int

	for _, char := range str {
		char := unicode.ToLower(char)

		if unicode.IsLetter(char) {
			switch char {
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
			default:
				consonants++
			}
		}
	}

	fmt.Printf("43. Vowels: %d\n", vowels)
	fmt.Printf("43. Consonants: %d\n", consonants)
}

func freqencyOfChar(str string) {
	var freq map[rune]int = make(map[rune]int)

	for _, char := range str {
		freq[char]++
	}

	fmt.Print("44. Freq of chars: ")
	for char, count := range freq {
		fmt.Printf("'%c': %d  ", char, count)
	}
	fmt.Println()
}

func removeSpaces(str string) {
	var i int = 0
	var res []rune = []rune(str)

	for _, char := range str {
		if char != ' ' {
			res[i] = char
			i++
		}
	}
	str = string(res[:i])

	fmt.Printf("45. Removed spaces: %s\n", str)
}
