package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, maxLength := 0, 1

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   // Palíndromos de comprimento ímpar
		len2 := expandAroundCenter(s, i, i+1) // Palíndromos de comprimento par
		length := maximum(len1, len2)

		if length > maxLength {
			start = i - (length-1)/2
			maxLength = length
		}
	}

	return s[start : start+maxLength]
}

// Expande em torno de um centro dado e retorna o comprimento do palíndromo.
func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// Retorna o máximo entre dois inteiros.
func maximum(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(longestPalindrome("babad")) // "bab" ou "aba"
	fmt.Println(longestPalindrome("cbbd"))  // "bb"
}
