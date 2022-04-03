package main

/*
 给定一个链表，删除链表的倒数第 n 个节点并返回链表的头指针
例如，
给出的链表为: 1→2→3→4→51\to 2\to 3\to 4\to 51→2→3→4→5, n=2n= 2n=2.
删除了链表的倒数第 nnn 个节点之后,链表变为1→2→3→51\to 2\to 3\to 51→2→3→5.

数据范围： 链表长度 0≤n≤10000\le n \le 10000≤n≤1000，链表中任意节点的值满足 0≤val≤1000 \le val \le 1000≤val≤100
要求：空间复杂度 O(1)O(1)O(1)，时间复杂度 O(n)O(n)O(n)
备注：
题目保证 nnn 一定是有效的
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	var rst *ListNode
	//返回倒数k-1 k个节点的地址
	//倒数第k-1有可能是nil
	jth, kth := findkthNode(head, n)
	if jth == nil {
		rst = kth
	} else {
		jth.Next = jth.Next.Next
		rst = head
	}
	return rst
}

func findkthNode(head *ListNode, k int) (*ListNode, *ListNode) {
	cnt := 0
	ln, rn := head, head
	for rn.Next != nil {
		rn = rn.Next
		cnt++
		if cnt > k {
			ln = ln.Next
		}
	}
	//只有一个节点的情况,肯定要删除一个的
	if cnt == 0 {
		return nil, nil
	}
	//如果第一个节点刚好是要删除的那一个
	if cnt == k-1 {
		return nil, head.Next
	}
	return ln, ln.Next
}
