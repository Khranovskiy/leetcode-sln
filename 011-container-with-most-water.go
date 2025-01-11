// 011. container-with-most-water
// https://leetcode.com/problems/container-with-most-water/
package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(h []int) int {
	var l, r = 0, len(h) - 1
	var res = int(0)

	for l < r {
		length := r - l
		minHeight := h[l]

		if h[r] < h[l] {
			minHeight = h[r]
		}

		square := length * minHeight

		if square > res {
			res = square
		}
		if h[l] < h[r] {
			l = l + 1
			continue
		}
		if h[r] < h[l] {
			r = r - 1
			continue
		}
		r = r - 1
		l = l + 1
	}
	return res
}
