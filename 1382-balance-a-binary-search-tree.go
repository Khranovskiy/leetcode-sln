package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
## Approach

1. **In-order Traversal**: Perform an in-order traversal of the BST to get all values in sorted order
2. **Build Balanced Tree**: Construct a balanced BST from the sorted array by:
  - Choosing the middle element as the root
  - Recursively building left subtree from left half
  - Recursively building right subtree from right half

## Complexity

- **Time**: O(n) - visit each node once for traversal, once for building
- **Space**: O(n) - array to store values plus recursion stack
*/

func balanceBST(root *TreeNode) *TreeNode {
	// Step 1: In-order traversal to get sorted vals
	var vals []int
	var inorderf func(*TreeNode)
	inorderf = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorderf(node.Left)
		vals = append(vals, node.Val)
		inorderf(node.Right)
	}
	inorderf(root)

	// Step 2: Build balanced BST from sorted array
	var bb func(int, int) *TreeNode
	bb = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := l + (r-l)/2
		n := &TreeNode{Val: vals[mid]}
		n.Left = bb(l, mid-1)
		n.Right = bb(mid+1, r)
		return n
	}

	return bb(0, len(vals)-1)
}

// Parse string like "[1,null,2,null,3,null,4]" and build tree
func buildTreeFromString(s string) *TreeNode {
	// Remove brackets and whitespace
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")

	if s == "" {
		return nil
	}

	// Split by comma
	parts := strings.Split(s, ",")

	// Convert to []*int
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

// Build tree from []*int array
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

		// Left child
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Helper to print tree structure
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

// Helper to print tree in-order
func printInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printInOrder(root.Left)
	fmt.Print(root.Val, " ")
	printInOrder(root.Right)
}

func main() {
	// Test case 1: [1,null,2,null,3,null,4]
	fmt.Println("Test Case 1:")
	fmt.Println("Original tree:")
	tree1 := buildTreeFromString("[1,null,2,null,3,null,4]")
	printTree(tree1, "", false)

	fmt.Println("\nBalanced tree:")
	balanced1 := balanceBST(tree1)
	printTree(balanced1, "", false)

	fmt.Print("\nIn-order: ")
	printInOrder(balanced1)
	fmt.Println("\n")

	// Test case 2: [2,1,3]
	fmt.Println("Test Case 2:")
	fmt.Println("Original tree:")
	tree2 := buildTreeFromString("[2,1,3]")
	printTree(tree2, "", false)

	fmt.Println("\nBalanced tree:")
	balanced2 := balanceBST(tree2)
	printTree(balanced2, "", false)

	fmt.Print("\nIn-order: ")
	printInOrder(balanced2)
	fmt.Println("\n")

	// Test case 3: [1,null,2,null,3,null,4,null,5,null,6,null,7]
	fmt.Println("Test Case 3:")
	fmt.Println("Original tree:")
	tree3 := buildTreeFromString("[1,null,2,null,3,null,4,null,5,null,6,null,7]")
	printTree(tree3, "", false)

	fmt.Println("\nBalanced tree:")
	balanced3 := balanceBST(tree3)
	printTree(balanced3, "", false)

	fmt.Print("\nIn-order: ")
	printInOrder(balanced3)
	fmt.Println()

	// Test case 4: Already balanced tree
	fmt.Println("Test Case 4:")
	fmt.Println("Original tree:")
	tree4 := buildTreeFromString("[4,2,6,1,3,5,7]")
	printTree(tree4, "", false)

	fmt.Println("\nBalanced tree:")
	balanced4 := balanceBST(tree4)
	printTree(balanced4, "", false)

	fmt.Print("\nIn-order: ")
	printInOrder(balanced4)
	fmt.Println()
}

/*
Test Case 1:
Original tree:
└── 1
    └── 2
        └── 3
            └── 4

Balanced tree:
└── 2
    ├── 1
    └── 3
        └── 4

In-order: 1 2 3 4

Test Case 2:
Original tree:
└── 2
    ├── 1
    └── 3

Balanced tree:
└── 2
    ├── 1
    └── 3

In-order: 1 2 3

Test Case 3:
Original tree:
└── 1
    └── 2
        └── 3
            └── 4
                └── 5
                    └── 6
                        └── 7

Balanced tree:
└── 4
    ├── 2
    │   ├── 1
    │   └── 3
    └── 6
        ├── 5
        └── 7

In-order: 1 2 3 4 5 6 7
Test Case 4:
Original tree:
└── 4
    ├── 2
    │   ├── 1
    │   └── 3
    └── 6
        ├── 5
        └── 7

Balanced tree:
└── 4
    ├── 2
    │   ├── 1
    │   └── 3
    └── 6
        ├── 5
        └── 7

In-order: 1 2 3 4 5 6 7
*/
