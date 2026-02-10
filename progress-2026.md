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

- m 3721. Longest Balanced Subarray 2 ❌
[url](https://leetcode.com/problems/longest-balanced-subarray-ii/description/)

??? min ❌✅

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^5

```go
func longestBalanced(nums []int) int {
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