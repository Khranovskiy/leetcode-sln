package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphaNumeric(rune(s[left])) {
			left++
		}
		for left < right && !isAlphaNumeric(rune(s[right])) {
			right--
		}
		if left < right && !equalsIgnoreCase(rune(s[left]), rune(s[right])) {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphaNumeric(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func equalsIgnoreCase(a, b rune) bool {
	return unicode.ToLower(a) == unicode.ToLower(b)
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(""))                               // true
}
