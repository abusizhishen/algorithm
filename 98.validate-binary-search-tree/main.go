package main

import "math"

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil{
		return true
	}

	return isBST(root,math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode,min,max int) bool {
	if root == nil{
		return true
	}

	if root.Val >= max || root.Val<=min{
		return false
	}

	return isBST(root.Left,min, root.Val) && isBST(root.Right, root.Val, max)
}