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