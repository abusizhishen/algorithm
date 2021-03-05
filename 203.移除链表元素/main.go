package main


 //Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil{
		return nil
	}

	head.Next = removeElements(head.Next, val)
	if head.Val == val{
		return head.Next
	}
	return head
}

//先删除头部的value的节点
//在下面只需要处理head.Next就好了
func removeElements2(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val{
		head = head.Next
	}

	p := head
	for p != nil && p.Next != nil {
		if p.Next.Val == val{
			p.Next = p.Next.Next
		}else{
			p = p.Next
		}
	}

	return head
}