package main

import (
	"fmt"
	"log"
)

type Node struct {
	Value int
	Next *Node
}

func reverse(node *Node) *Node{
	var pre *Node
	var head = node
	for head != nil{
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}

func huiwen(node *Node)bool  {
	if node == nil{
		return false
	}

	middle := findMiddle(node)
	middle=reverse(middle)
	log.Println("find middle")
	printNode(middle)
	for middle != nil {
		if middle.Value != node.Value{
			return false
		}

		middle = middle.Next
		node = node.Next
	}

	return true
}

func nodeLength(head *Node)int  {
	if head == nil{
		return 0
	}

	var i int
	node := head
	for node != nil{
		i++
		node = node.Next
	}

	return i
}

func main()  {
	var arr = []int{1,2,3,4,5,4,3,2,1}
	var node = construct(arr)
	printNode(node)
	result := huiwen(node)
	fmt.Println(result)
}

func construct(arr []int) *Node {
	if len(arr) == 0{
		return nil
	}

	var head = &Node{
		Value: arr[0],
		Next:  nil,
	}
	var result = head

	for i:=1;i<len(arr);i++{
		node := &Node{
			Value: arr[i],
			Next:  nil,
		}

		head.Next = node
		head = node
	}

	return result
}

func findMiddle(node *Node) *Node {
	if node == nil{
		return nil
	}

	var length = nodeLength(node)
	fmt.Println("len:", length)
	middle := length/2
	result := node
	for i:=0;i<=middle;i++{
		result = result.Next
	}

	return result
}

func printNode(node *Node)  {
	for node != nil{
		log.Printf("%d ", node.Value)
		node = node.Next
	}

	fmt.Println()
}