package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "  hello   world  "
	// s := "the sky   is blue"
	s := "a          b"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	// word separated at least one space
	// return string of the words in reverse order concatenated by a signle space

	sRunes := trimSpaces([]rune(s))
	lWb := 0
	rWb := 0
	for i := 0; i < len(sRunes); i++ {
		// skip leading spaces
		if sRunes[i] == ' ' && lWb == rWb {
			lWb++
			rWb++
		}
		if sRunes[i] == ' ' && lWb != rWb {
			reverseWord(sRunes, lWb, rWb)
			rWb += 1
			lWb = rWb
			continue

		}
		if sRunes[i] != ' ' && rWb < len(sRunes) {
			if sRunes[lWb] == ' ' {
				lWb = i
			}
			rWb = i
		}
	}
	if lWb != rWb {
		reverseWord(sRunes, lWb, rWb)
	}
	reverseWord(sRunes, 0, len(sRunes)-1)

	return strings.TrimSpace(string(sRunes))
}
func reverseWord(s []rune, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func trimSpaces(sRunes []rune) []rune {
	isStart := true
	j := 0
	for i := 0; i < len(sRunes); i++ {
		if sRunes[i] == ' ' && isStart {
			continue
		}

		sRunes[j] = sRunes[i]
		j++
		isStart = sRunes[i] == ' '
	}

	if sRunes[j-1] == ' ' {
		j--
	}

	return sRunes[:j]
}
