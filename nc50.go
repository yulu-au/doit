package main

/*
将给出的链表中的节点每 k 个一组翻转，返回翻转后的链表
如果链表中的节点数不是 k 的倍数，将最后剩下的节点保持原样
你不能更改节点中的值，只能更改节点本身。

数据范围：  0≤n≤2000\ 0 \le n \le 2000 0≤n≤2000 ， 1≤k≤20001 \le k \le 20001≤k≤2000 ，链表中每个元素都满足 0≤val≤10000 \le val \le 10000≤val≤1000
要求空间复杂度 O(1)O(1) O(1)，时间复杂度 O(n)O(n)O(n)
例如：
给定的链表是 1→2→3→4→51\to2\to3\to4\to51→2→3→4→5
对于 k=2k = 2k=2 , 你应该返回 2→1→4→3→52\to 1\to 4\to 3\to 52→1→4→3→5
对于 k=3k = 3k=3 , 你应该返回 3→2→1→4→53\to2 \to1 \to 4\to 53→2→1→4→5
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var subHead, rst *ListNode
	etyp := &ListNode{}
	cur := etyp.Next
	cnt := k
	for head != nil {
		if cnt == k {
			subHead = head
		}
		cnt--
		head = head.Next
		if cnt == 0 {
			nt := head.Next
			head.Next = nil

			cur.Next = ReverseList(subHead)
			cur = subHead
			cnt = k

			head = nt
		}
	}
	cur.Next = subHead

	rst = etyp.Next
	etyp.Next = nil
	return rst
}
