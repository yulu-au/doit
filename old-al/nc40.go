package main

/*
假设链表中每一个节点的值都在 0 - 9 之间，那么链表整体就可以代表一个整数。
给定两个这种链表，请生成代表两个整数相加值的结果链表。
数据范围：0≤n,m≤10000000 \le n,m \le 10000000≤n,m≤1000000，链表任意值 0≤val≤90 \le val \le 9 0≤val≤9
要求：空间复杂度 O(n)O(n)O(n)，时间复杂度 O(n)O(n)O(n)

例如：链表 1 为 9->3->7，链表 2 为 6->3，最后生成新的结果链表为 1->0->0->0。
*/

func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	reverseH1, reverseH2 := ReverseList(head1), ReverseList(head2)
	extra := 0
	ety := &ListNode{}
	cur := ety

	for reverseH1 != nil && reverseH2 != nil {
		sum := reverseH1.Val + reverseH2.Val + extra
		extra = sum / 10
		cur.Next = &ListNode{Val: sum % 10}

		cur = cur.Next
		reverseH1 = reverseH1.Next
		reverseH2 = reverseH2.Next
	}
	for reverseH1 != nil {
		sum := extra + reverseH1.Val
		extra = sum / 10
		cur.Next = &ListNode{Val: sum % 10}

		cur = cur.Next
		reverseH1 = reverseH1.Next
	}
	for reverseH2 != nil {
		sum := extra + reverseH2.Val
		extra = sum / 10
		cur.Next = &ListNode{Val: sum % 10}

		cur = cur.Next
		reverseH2 = reverseH2.Next
	}

	if extra > 0 {
		cur.Next = &ListNode{Val: extra}
		cur = cur.Next
	}
	res := ety.Next
	ety.Next = nil
	res = ReverseList(res)

	return res
}
