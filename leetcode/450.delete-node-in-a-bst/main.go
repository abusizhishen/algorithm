package main


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil{
		return nil
	}

	if root.Val > key{
		root.Left = deleteNode(root.Left,key)
		return root
	}

	if root.Val < key{
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Left == nil{
		return root.Right
	}

	if root.Right == nil{
		return root.Left
	}

	left := root.Left
	right := root.Right
	for right.Left != nil{
		right = right.Left
	}

	right.Left = left
	root.Left = nil
	right = root.Right
	root = nil

	return right
}