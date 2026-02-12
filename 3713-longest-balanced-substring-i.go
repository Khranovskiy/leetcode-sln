package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestBalanced3("abbac")) // 4
	fmt.Println(longestBalanced3("zzabccy"))
	fmt.Println(longestBalanced3("aba"))
}

func longestBalanced3(s string) int {
	ans := 0
	for i, _ := range s {

		var (
			freq     = make([]int, 26)
			maxFreq  = 0
			distinct = 0
		)

		for j := i; j < len(s); j++ {
			index := int(s[j]) - 'a'

			if freq[index] == 0 {
				distinct++
			}

			freq[index]++
			if freq[index] > maxFreq {
				maxFreq = freq[index]
			}

			length := j - i + 1

			if length == maxFreq*distinct {
				if length > ans {
					ans = length
				}
			}
		}
	}
	return ans
}
