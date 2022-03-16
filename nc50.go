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

/*
写这道题的时候遇到了很多边界的问题,体会是函数是复杂度的封装,如果你认为现在写起来很乱,
那么很明显,应当抽象出一个函数来向下转移复杂度了
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var rst *ListNode
	empty := &ListNode{}
	cur := empty

	//没有剩余链表就停止
	for head != nil {
		lReverse, lRemain := spiltAndReverse(head, k)
		cur.Next = lReverse
		cur = head
		head = lRemain
	}

	rst = empty.Next
	return rst
}

//分割并反转一部分链表,并且返回剩下的部分链表
func spiltAndReverse(l *ListNode, k int) (*ListNode, *ListNode) {
	cnt := k
	//反转的链表和剩余的
	var lReverse, lRemain *ListNode
	head := l
	for ; l != nil; l = l.Next {
		cnt--
		if cnt == 0 {
			if l.Next == nil {
				lRemain = nil
			} else {
				lRemain = l.Next
				//为翻转做准备,链表分成两部分
				l.Next = nil
			}
			lReverse = ReverseList(head)
			break
		}
	}
	//链表节点数目不足以翻转
	if cnt != 0 {
		lRemain = nil
		lReverse = head
	}

	return lReverse, lRemain
}
