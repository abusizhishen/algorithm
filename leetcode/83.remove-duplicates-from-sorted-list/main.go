package main


type ListNode struct {
  Val int
  Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	var h = head

	for h != nil && h.Next != nil{
		if h.Val != h.Next.Val{
			h = h.Next
			continue
		}

		h.Next = h.Next.Next
	}

	return head
}
