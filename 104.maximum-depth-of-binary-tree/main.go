package main

func maxDepth(tree *TreeNode) int {
	if tree == nil{
		return 0
	}
	return max(maxDepth(tree.Left),maxDepth(tree.Right))+1
}

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func max(a,b int) int {
	if a > b{
		return a
	}

	return b
}