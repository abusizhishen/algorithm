package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum = 1
	if root.Left != nil {
		sum += countNodes(root.Left)
	}

	if root.Right != nil {
		sum += countNodes(root.Right)
	}

	return sum
}
