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