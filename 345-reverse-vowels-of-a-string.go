package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
}

func reverseVowels(s string) string {
	lp, rp := 0, len(s)-1
	runes := []rune(s)
	for lp < rp {
		if !isVowel(runes[lp]) {
			lp++
		}
		if !isVowel(runes[rp]) {
			rp--
		}
		if isVowel(runes[lp]) && isVowel(runes[rp]) {
			runes[lp], runes[rp] = runes[rp], runes[lp]
			lp++
			rp--
		}
	}
	return string(runes)
}
func isVowel(ch rune) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
