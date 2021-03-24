package main

/**
Definition for a binary tree node.
*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


//层次遍历

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var nodes []*TreeNode
	nodes = append(nodes, root)

	var values []int
	for len(nodes) !=0 {
		var temp []*TreeNode
		for _,node := range nodes{
			if node.Left != nil {
				temp = append(temp, node.Left)
			}

			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}

		values = append(values, nodes[len(nodes)-1].Val)
		nodes = temp
	}

	return values
}
