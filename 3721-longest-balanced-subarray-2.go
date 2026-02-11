package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestBalanced2([]int{1, 2, 3, 2}))    // 3
	fmt.Println(longestBalanced2([]int{3, 2, 2, 5, 4})) // 5

	fmt.Println(longestBalanced2([]int{3}))    // expected 0
	fmt.Println(longestBalanced2([]int{6, 2})) // expected 0

	fmt.Println(longestBalanced2([]int{15, 6, 15, 11})) // expected 3
}

// LeetCode 3721. Longest Balanced Subarray II
// O(n log n) time, O(n) space
// Technique: Segment Tree with Lazy Propagation + Prefix Sum + Last Occurrence trick
//
// IDEA:
// For each right endpoint j, maintain PS[i] = (distinctEvens - distinctOdds)
// for every subarray [i, j]. A subarray is balanced when PS[k] == PS[j+1].
//
// When processing nums[j]:
//   - If nums[j] appeared before at index p, subarrays starting at i ≤ p
//     already counted it. Undo: range add -contrib to PS[p+1..n].
//   - New contribution: range add +contrib to PS[j+1..n].
//   - Find leftmost k in [0, j] where PS[k] == PS[j+1] → answer candidate j-k+1.
//
// The "find leftmost position with value == target" query is O(log n) because
// consecutive PS values differ by at most 1 (continuity), enabling tight min/max pruning.

func longestBalanced2(nums []int) int {
	n := len(nums)
	sz := n + 1 // segment tree positions 0..n

	mn := make([]int, 4*sz)
	mx := make([]int, 4*sz)
	lz := make([]int, 4*sz)

	pushDown := func(o int) {
		if lz[o] != 0 {
			for _, c := range []int{o * 2, o*2 + 1} {
				mn[c] += lz[o]
				mx[c] += lz[o]
				lz[c] += lz[o]
			}
			lz[o] = 0
		}
	}

	pushUp := func(o int) {
		mn[o] = min(mn[o*2], mn[o*2+1])
		mx[o] = max(mx[o*2], mx[o*2+1])
	}

	var update func(o, l, r, ql, qr, val int)
	update = func(o, l, r, ql, qr, val int) {
		if ql > qr || l > qr || r < ql {
			return
		}
		if ql <= l && r <= qr {
			mn[o] += val
			mx[o] += val
			lz[o] += val
			return
		}
		pushDown(o)
		mid := (l + r) / 2
		update(o*2, l, mid, ql, qr, val)
		update(o*2+1, mid+1, r, ql, qr, val)
		pushUp(o)
	}

	var pointQuery func(o, l, r, pos int) int
	pointQuery = func(o, l, r, pos int) int {
		if l == r {
			return mn[o]
		}
		pushDown(o)
		mid := (l + r) / 2
		if pos <= mid {
			return pointQuery(o*2, l, mid, pos)
		}
		return pointQuery(o*2+1, mid+1, r, pos)
	}

	var findLeftmost func(o, l, r, ql, qr, target int) int
	findLeftmost = func(o, l, r, ql, qr, target int) int {
		if ql > qr || l > qr || r < ql {
			return -1
		}
		if mn[o] > target || mx[o] < target {
			return -1
		}
		if l == r {
			return l
		}
		pushDown(o)
		mid := (l + r) / 2
		res := findLeftmost(o*2, l, mid, ql, qr, target)
		if res != -1 {
			return res
		}
		return findLeftmost(o*2+1, mid+1, r, ql, qr, target)
	}

	prev := make(map[int]int) // value → last seen index
	ans := 0
	N := sz - 1 // max segment tree index = n

	for j := 0; j < n; j++ {
		v := nums[j]
		contrib := 1
		if v%2 == 1 {
			contrib = -1
		}

		if p, ok := prev[v]; ok {
			update(1, 0, N, p+1, N, -contrib)
		}
		update(1, 0, N, j+1, N, contrib)
		prev[v] = j

		target := pointQuery(1, 0, N, j+1)
		k := findLeftmost(1, 0, N, 0, j, target)
		if k != -1 {
			if length := j - k + 1; length > ans {
				ans = length
			}
		}
	}

	return ans
}
