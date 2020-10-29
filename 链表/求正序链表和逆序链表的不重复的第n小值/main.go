package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

func reverse(node *Node) *Node {
	if node == nil{
		return node
	}

	var pre *Node
	var temp = node
	for temp != nil{
		next := temp.Next
		temp.Next = pre
		pre = temp
		temp = next
	}

	return pre
}

func merge(a,c *Node, k int) int {
	var i,min,numa,numc int
	for i<k{
		for a !=nil   && a.Next!= nil{
			if a.Value != a.Next.Value{
				break
			}
			a = a.Next
		}

		numa = a.Value
		for c !=nil && c.Next != nil{
			if c.Value != c.Next.Value{
				break
			}

			c = c.Next
		}
		numc = c.Value

		min = numa
		if numa>numc{
			min = numc
			c = c.Next
		}else if numa <numc{
			a = a.Next
		}else {
			a = a.Next
			c = c.Next
		}

		i++
		fmt.Printf("i:%d min:%d numa:%d numc:%d \n", i,min,numa,numc)


		if i == k{
			return min
		}

	}

	return min
}

func main()  {
	var a = []int{10,8,7,5,3}
	var b = []int{1,3,5,7,9}

	var nodeA = construct(a)
	var nodeB = construct(b)
	nodeB = reverse(nodeB)
	print(nodeB)
	print(nodeA)
	var num = merge(nodeA,nodeB, 5)
	fmt.Println(num)
}

func construct(arr []int) *Node {
	var node *Node
	for _,num := range arr{
		var temp = &Node{
			Value: num,
			Next:  node,
		}

		node = temp
	}

	return node
}

func print(node *Node)  {
	for node != nil{
		fmt.Print(node.Value," ")
		node = node.Next
	}
	fmt.Println()
}

