package main

import "fmt"

//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null
//Definition for singly-linked list.
  type ListNode struct {
      Val int
      Next *ListNode
  }

  //假设链表环前有 aa 个节点，环内有 bb 个节点
//本题核心思路：走 a+nba+nb 步一定处于环的入口位置
//
//利用快慢指针 fastfast 和 slowslow，fastfast 一次走两步，slowslow 一次走一步
//当两个指针第一次相遇时，假设 slowslow 走了 ss 步，下面计算 fastfast 走过的步数
//i. fastfast 比 slowslow 多走了 nn 个环：f = s + nbf=s+nb
//ii. fastfast 比 slowslow 多走一倍的步数：f = 2sf=2s --> 跟上式联立可得 s = nbs=nb
//iii. 综上计算得，f = 2nbf=2nb，s = nbs=nb
//也就是两个指针第一次相遇时，都走过了环的倍数，那么再走 aa 步就可以到达环的入口
//让 fastfast 从头再走，slowslow 留在原地，fastfast 和 slowslow 均一次走一步，当两个指针第二次相遇时，fastfast 走了 aa 步，slowslow 走了 a+nba+nb 步
//此时 slowslow 就在环的入口处，返回 slowslow
//
//作者：edelweisskoko
//链接：https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/142-huan-xing-lian-biao-iishuang-zhi-zhe-k8ju/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func detectCycle(head *ListNode) *ListNode {
	fast,slow := head,head

	for fast != nil && fast.Next != nil{
		//todo 先赋值再比较
		fast = fast.Next.Next
		slow = slow.Next
		if fast != slow {
			continue
		}

		res := head
		for slow != res{
			slow = slow.Next
			res = res.Next
		}

		return res
	}

	return nil
}

func main() {
	var a = ListNode{Val:  4, Next: nil,}
	var b = ListNode{Val:  0, Next: nil,}
	var c = ListNode{Val:  2, Next: nil,}
	var d = ListNode{Val:  3, Next: nil,}

	d.Next = &c
	c.Next = &b
	b.Next = &a
	a.Next = &c

	fmt.Println(detectCycle(&d))
}