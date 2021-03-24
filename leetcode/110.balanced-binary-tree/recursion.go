package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	return jian(height(root.Left), height(root.Right))
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return height(root.Right) + 1
	}

	if root.Right == nil {
		return height(root.Left) + 1
	}

	return max(height(root.Left), height(root.Right)) + 1
}

func jian(a, b int) bool {
	if a > b {
		b, a = a, b
	}

	return b-a <= 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
