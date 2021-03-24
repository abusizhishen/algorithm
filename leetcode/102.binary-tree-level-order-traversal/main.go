package main


//  Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func levelOrder(root *TreeNode) [][]int {
	if root == nil{
		return nil
	}
	var result = [][]int{[]int{root.Val}}
	var arr = []*TreeNode{root}

	for len(arr) !=0{
		var tmp []*TreeNode
		var ints []int
		for _,node := range arr{
			if node.Left != nil{
				ints = append(ints,node.Left.Val)
				tmp = append(tmp, node.Left)
			}

			if node.Right != nil{
				ints = append(ints,node.Right.Val)
				tmp = append(tmp, node.Right)
			}
		}
		if len(ints) > 0{
			result = append(result, ints)
		}

		arr = tmp
	}

	return result
}
