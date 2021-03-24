package main

type ListNode struct {
      Val int
      Next *ListNode
  }

func hasCycle(head *ListNode) bool {
	var h1,h2 = head,head

	for h1 != nil && h2 != nil && h2.Next!= nil{
		h1 = h1.Next
		h2 = h2.Next.Next
		if h1 == h2 {
			return true
		}
	}
	return false
}
