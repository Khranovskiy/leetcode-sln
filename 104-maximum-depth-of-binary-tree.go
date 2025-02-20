package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, level int) int {

	// Если в функции dfs на вход пришла пустая нода,
	// то это означает, что мы опустились ниже последнего листового узла в этой ветке,
	// т. е. мы спустились до ее конца.
	// В этом случае мы возвращаем из функции level в качестве ответа,
	// так как он обозначает текущую глубину ветки.

	if node == nil {
		return level
	}

	l := dfs(node.Left, level+1)
	r := dfs(node.Right, level+1)

	if l > r {
		return l
	}
	return r
}
