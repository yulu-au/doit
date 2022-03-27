package main

/*
判断给定的链表中是否有环。如果有环则返回true，否则返回false。
*/

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	a, b := head, head

	for a != nil && b != nil {
		a = a.Next
		b = b.Next
		if b != nil {
			b = b.Next
		}
		//如果是无环链表,这里a可能已经到链表结尾
		if a == b && a != nil {
			return true
		}
	}
	return false
}
