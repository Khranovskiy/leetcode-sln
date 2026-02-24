progress-2026.md

=================================================================

- m 1382. Balance a Binary Search Tree 
[1382. Balance a Binary Search Tree](https://leetcode.com/problems/balance-a-binary-search-tree/description/?envType=daily-question&envId=2026-02-09)

15 min

```go
func balanceBST(root *TreeNode) *TreeNode {
	// Step 1: In-order traversal to get sorted values
	var values []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		values = append(values, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	// Step 2: Build balanced BST from sorted array
	var bb func(int, int) *TreeNode
	bb = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := l + (r-l)/2
		n := &TreeNode{Val: values[mid]}
		n.Left = bb(l, mid-1)
		n.Right = bb(mid+1, r)
		return n
	}

	return bb(0, len(values)-1)
}
```

=================================================================

- m 3719. Longest Balanced Subarray I ❌
[url](https://leetcode.com/problems/longest-balanced-subarray-i/description/?envType=daily-question&envId=2026-02-10)

20 min ❌
next 35 min - wrong answer ❌ 

1 <= nums.length <= 1500
1 <= nums[i] <= 10^5

```go
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
```

=================================================================

- H 3721. Longest Balanced Subarray 2 ❌
[url](https://leetcode.com/problems/longest-balanced-subarray-ii/description/)

??? min ❌✅

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^5

https://claude.ai/share/ee1f0a77-9bdf-4984-84a7-6287ad845605

```go
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

func longestBalanced(nums []int) int {
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
```
=================================================================

- m 3713. Longest Balanced Substring I
[url](https://leetcode.com/problems/longest-balanced-substring-i/description/?envType=daily-question&envId=2026-02-12)

bruteforce solution due to small input data

Constraints:
- 1 <= s.length <= 1000
- s consists of lowercase English letters.

inf min ❌

- frequency array of size 26
- distinct → how many unique characters
- maxFreq → maximum frequency in current substring

Runtime 19 ms Beats 100.00%
Memory 4.58 MB Beats 100.00%


```go
func longestBalanced(s string) int {
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
```


=================================================================

- e 67. Add Binary
[url](https://leetcode.com/problems/add-binary)

10 min ❌

```go
func addBinary(a string, b string) string {
	result := make([]byte, max(len(a), len(b))+1)

	i, j := len(a)-1, len(b)-1
	k := len(result) - 1
	carry := byte(0)

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i>=0 {sum += a[i] - '0';i--}
		if j>=0 {sum += b[j]- '0'; j--}

		result[k] = (sum % 2) + '0'
		carry = sum / 2
		k--
	}
	for k >=0 {
		result[k]  = '0'
		k--
	}
	start :=0
	for start < len(result)-1 && result[start] == '0'{
		start++
	}

	return string(result[start:])
}
```

=================================================================

- e 190. Reverse Bits
[url](https://leetcode.com/problems/reverse-bits/description/?envType=daily-question&envId=2026-02-16)

??? min ❌ ✅

```go
// For each of the 32 iterations:

// result << 1 — shift the result left to make room for the next bit.
// num & 1 — extract the lowest bit of num.
// | — place that bit into the lowest position of result.
// num >>= 1 — shift num right to process the next bit.


func reverseBitsSlow(num int) int {
	var result int
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num & 1)
		num >>= 1
	}
	return result
}

var cache [256]int

func initBetter() {
	for i := 0; i < 256; i++ {
		var rev int
		val := i
		for j := 0; j < 8; j++ {
			rev = (rev << 1) | (val & 1)
			val >>= 1
		}
		cache[i] = rev
	}
}

func reverseBitsBetter(num int) int {
	return cache[num&0xff]<<24 |
		cache[(num>>8)&0xff]<<16 |
		cache[(num>>16)&0xff]<<8 |
		cache[(num>>24)&0xff]
}
// best
func reverseBits(num int) int {
    num = (num>>16)&0x0000FFFF | (num&0x0000FFFF)<<16
    num = (num>>8)&0x00FF00FF | (num&0x00FF00FF)<<8
    num = (num>>4)&0x0F0F0F0F | (num&0x0F0F0F0F)<<4
    num = (num>>2)&0x33333333 | (num&0x33333333)<<2
    num = (num>>1)&0x55555555 | (num&0x55555555)<<1
    return num
}
```
=================================================================
- e 401. Binary Watch

[url](https://leetcode.com/problems/binary-watch/description/?envType=daily-question&envId=2026-02-17)

??? min ✅

```go

func readBinaryWatch(turnedOn int) []string {
	// 0-11.             1 0 1 1 -> 0..3 leds
	// 0-59.    0 0 1 1  1 0 1 1 -> 0..5 leds

	var result = make([]string, 0)

	for h := 0; h <= 11; h++ {
		for m:=0; m <= 59; m++ {
			if countBits(h) + countBits(m) == turnedOn {
				result = append(result, fmt.Sprintf("%01d:%02d", h, m))
			}
		}
	}
	return result
}

func countBits(n int) int {
	var res int
	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}
```


693. Binary Number with Alternating Bits
easy

https://leetcode.com/problems/binary-number-with-alternating-bits/?envType=daily-question&envId=2026-02-18

func hasAlternatingBits(n int) bool {
    x := n ^ (n >> 1)
    return x & (x + 1) == 0
}


https://leetcode.com/problems/special-binary-string/editorial/?envType=daily-question&envId=2026-02-20
h 761. Special Binary String

```go
func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	var parts []string
	balance := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			balance++
		} else {
			balance--
		}
		// Found a top-level special substring s[start : i+1]
		if balance == 0 {
			// Strip outer '1' and '0', recursively optimize inside
			inner := makeLargestSpecial(s[start+1 : i])
			parts = append(parts, "1"+inner+"0")
			start = i + 1
		}
	}

	// Put larger blocks first
	sort.Slice(parts, func(i, j int) bool {
		return parts[i] > parts[j]
	})

	return strings.Join(parts, "")
}
```


https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/description/?envType=daily-question&envId=2026-02-21

e 762. Prime Number of Set Bits in Binary Representation


```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(97)
	result := n.ProbablyPrime(20) // 20 = number of rounds
	fmt.Println(result) // true

	fmt.Println(countPrimeSetBits(10, 15))
}

func countPrimeSetBits(left int, right int) int {
	result := 0

	for i := left; i <= right; i++ {
		bits := countBits(i)
		isPrime := big.NewInt(int64(bits)).ProbablyPrime(10)
		if isPrime {
			result++
		}
	}
	return result
}

func countBits(n int) int {
	var res int
	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}
```


M 1461. Check If a String Contains All Binary Codes of Size K
https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/?source=submission-ac

1461-check-if-a-string-contains-all-binary-codes-of-size-k.go
```go
func hasAllCodes(s string, k int) bool {
	table := make(map[int64]struct{})
	if k >= len(s) {
		return false
	}

	for l, r := 0, k; r <= len(s); {
		substring := s[l:r]
		l++;r++

		key,_ := strconv.ParseInt(substring, 2, 64)
		table[key] = struct{}{}
	}

	for i := int64(0); i < (2 << (k - 1)); i++ {
		if _, ok := table[i]; !ok {
			return false
		}
	}
	return true
}
```

e 1022. Sum of Root To Leaf Binary Numbers
https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/?envType=daily-question&envId=2026-02-24

1022-sum-of-root-to-leaf-binary-numbers.go

```go
func sumRootToLeaf(root *TreeNode) int {
    var dfs func(node *TreeNode, cur int) int
    dfs = func(node *TreeNode, cur int) int {
        if node == nil {
            return 0
        }
        cur = cur*2 + node.Val
        if node.Left == nil && node.Right == nil {
            return cur
        }
        return dfs(node.Left, cur) + dfs(node.Right, cur)
    }
    return dfs(root, 0)
}
```

=================================================================

# template 

=================================================================
- m 
[url]( )

??? min ❌ ✅

```go
```

=================================================================

# utils

=================================================================

## Print BST

```go
// Helper to print tree structure
// 
// Balanced tree:
// └── 2
//     ├── 1
//     └── 3
func printTree(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	connector := "└── "
	if isLeft {
		connector = "├── "
	}
	fmt.Println(prefix + connector + fmt.Sprint(root.Val))

	if root.Left != nil || root.Right != nil {
		extension := "    "
		if isLeft {
			extension = "│   "
		}
		if root.Left != nil {
			printTree(root.Left, prefix+extension, true)
		}
		if root.Right != nil {
			printTree(root.Right, prefix+extension, false)
		}
	}
}
```

```go
// Helper to print tree in-order
// 
// In-order: 1 2 3 
func printInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printInOrder(root.Left)
	fmt.Print(root.Val, " ")
	printInOrder(root.Right)
}
```
## Build BST from string

Parse string like "[1,null,2,null,3,null,4]" and build tree

usage

```go
	tree1 := buildTreeFromString("[1,null,2,null,3,null,4]")
	printTree(tree1, "", false)
```

```go
func buildTreeFromString(s string) *TreeNode {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")

	if s == "" {
		return nil
	}

	parts := strings.Split(s, ",")

	values := make([]*int, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "null" {
			values = append(values, nil)
		} else {
			val, err := strconv.Atoi(part)
			if err != nil {
				panic(fmt.Sprintf("Invalid number: %s", part))
			}
			values = append(values, &val)
		}
	}

	return buildTree(values)
}

func buildTree(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
```