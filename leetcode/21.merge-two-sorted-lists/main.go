package main

//Definition for singly-linked list.
 type ListNode struct {
      Val int
      Next *ListNode
 }
 
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var l3 = &ListNode{}
    var l4 = l3
    for l1 != nil && l2 != nil{
        var node ListNode
        if l1.Val < l2.Val{
           node.Val = l1.Val
           l1 = l1.Next
        }else{
            node.Val = l2.Val
            l2 = l2.Next
        }

        l3.Next = &node
        l3 = &node
    } 

    if l1 == nil{
        l3.Next = l2
    }else{
        l3.Next = l1
    }

    return l4.Next
}
