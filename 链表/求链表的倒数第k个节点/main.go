package main

import "fmt"

//输入一个链表，输出该链表中倒数第k个节点。从1开始计数，即链表的尾节点是倒数第1个节点。
//例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。
//示例：
//给定一个链表: 1->2->3->4->5, 和 k = 2.
//返回链表 4->5.


func main()  {
	var node = &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next:  &Node{
				Value: 3,
				Next:  &Node{
					Value: 4,
					Next:  &Node{
						Value: 5,
						Next:  nil,
					},
				},
			},
		},
	}

	print(node)
	print(output(node,2))
}

type Node struct{
	Value int
	Next *Node
}

func output(node *Node,k int) *Node {
	var p1,p2 = node,node

	for i:=0;i<k;i++{
		p1 = p1.Next
	}

	for p1 != nil{
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}

//func constructNode(arr []int) *Node {
//
//}

func print(node *Node)  {
	for node != nil{
		fmt.Print(node.Value)
		node = node.Next
	}

	fmt.Println()
}