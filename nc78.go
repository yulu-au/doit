package main

/*
给定一个单链表的头结点pHead(该头节点是有值的，比如在下图，它的val是1)，长度为n，反转该链表后，返回新链表的表头
*/

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}

	cur := pHead
	pFirst := pHead

	for cur.Next != nil {
		//1 remove a node after cur from the list
		tmp := cur.Next
		cur.Next = cur.Next.Next
		//1
		//2 update head of the list
		tmp.Next = pFirst
		pFirst = tmp
	}

	return pFirst
}
