package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(decodeString("3[a2[c]]")) // accaccacc
}

func decodeString(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			res += string(s[i])
			continue
		}

		count := 0
		for s[i] != '[' {
			count = count*10 + int(s[i]-'0')
			i++
		}
		bracket := 1
		j := i + 1
		for bracket > 0 {
			if s[j] == '[' {
				bracket++
			}
			if s[j] == ']' {
				bracket--
			}
			j++
		}
		subStr := s[i+1 : j-1]
		inner := decodeString(subStr)
		res += strings.Repeat(inner, count)
		i = j - 1
	}

	return res
}
