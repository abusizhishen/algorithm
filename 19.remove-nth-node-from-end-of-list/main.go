package main

type ListNode struct {
      Val int
      Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var h = &ListNode{Next:head}
	var h2,h1 = h,h

	for i:=0;i<=n;i++{
		h1 = h1.Next
	}

	for h1!=nil{
		h1 = h1.Next
		h2 = h2.Next
	}


	h2.Next = h2.Next.Next
	return h.Next
}
