package main

/*
 输入两个递增的链表，单个链表的长度为n，合并这两个链表并使新链表中的节点仍然是递增排序的。
数据范围： 0≤n≤10000 \le n \le 10000≤n≤1000，−1000≤节点值≤1000-1000 \le 节点值 \le 1000−1000≤节点值≤1000
要求：空间复杂度 O(1)O(1)O(1)，时间复杂度 O(n)O(n)O(n)

如输入{1,3,5},{2,4,6}时，合并后的链表为{1,2,3,4,5,6}，所以对应的输出为{1,2,3,4,5,6}
*/
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}

	var resHead, cur *ListNode
	if pHead1.Val < pHead2.Val {
		resHead = pHead1
		cur = pHead1
		pHead1 = pHead1.Next
	} else {
		resHead = pHead2
		cur = pHead2
		pHead2 = pHead2.Next
	}

	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		}
		cur = cur.Next
	}

	if pHead1 == nil {
		cur.Next = pHead2
	} else {
		cur.Next = pHead1
	}

	return resHead
}
