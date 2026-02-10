package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(longestBalanced([]int{1, 2, 3, 2}))    // 3
	fmt.Println(longestBalanced([]int{3, 2, 2, 5, 4})) // 5

	fmt.Println(longestBalanced([]int{3}))    // expected 0
	fmt.Println(longestBalanced([]int{6, 2})) // expected 0

	fmt.Println(longestBalanced([]int{15, 6, 15, 11})) // expected 3
}

// 23:30
// 23:49 first attempt
// 00:04 wrong answer- test case [6,2]

func longestBalanced(nums []int) int {
	maxLen := 0

	for i := 0; i < len(nums); i++ {
		odd := make(map[int]int)
		even := make(map[int]int)

		for j := i; j < len(nums); j++ {
			if nums[j]&1 == 1 {
				odd[nums[j]]++
			} else {
				even[nums[j]]++
			}
			if len(odd) == len(even) {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
				}
			}
		}
	}
	return maxLen
}

func longestBalancedWrong(ar []int) int {
	var (
		counts      map[int]int = make(map[int]int)
		evens       int
		odds        int
		left, right int = 0, len(ar) - 1
	)
	for _, v := range ar {

		if _, exists := counts[v]; !exists {
			counts[v] = 1

			if isOdd(v) {
				odds += 1
			} else {
				evens += 1
			}
		} else {
			counts[v] = counts[v] + 1
		}
	}
	for left < right && odds != evens {
		var n int = 0
		if odds > evens {
			if isOdd(ar[left]) {
				n = ar[left]
				left += 1
			} else if isOdd(ar[right]) {
				n = ar[right]
				right -= 1
			} else {
				continue
			}
			counts[n]--
			if counts[n] == 0 {
				odds--
			}
		}
		if evens > odds {
			if isEven(ar[left]) {
				n = ar[left]
				left += 1
			} else if isEven(ar[right]) {
				n = ar[right]
				right -= 1
			} else {
				continue
			}
			counts[n]--
			if counts[n] == 0 {
				evens--
			}
		}
	}
	if odds == 0 && evens != odds {
		return 0
	}
	if evens == 0 && evens != odds {
		return 0
	}
	return right - left + 1
}

func isOdd(n int) bool {
	return n%2 == 1
}
func isEven(n int) bool {
	return n%2 == 0
}

func longestBalancedTest(t *testing.T) {
	fmt.Println(longestBalanced([]int{1, 2, 3, 2}))
}
