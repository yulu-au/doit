package main

/*
输入两个无环的单向链表，找出它们的第一个公共结点，如果没有公共节点则返回空。（注意因为传入数据是链表，所以错误测试数据的提示是用其他方式显示的，保证传入数据是正确的）

数据范围： n≤1000n \le 1000n≤1000
要求：空间复杂度 O(1)O(1)O(1)，时间复杂度 O(n)O(n)O(n)

*/

/*
如果两个链表长度相等且有相交,那找到这个相交点是很容易的.两个指针同时走,相等的时候就是相交点.
还有链表长度不等的情况,这种情况可以转化为上面相等的情况.计算出两个链表的长度,然后虚拟对齐
*/
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	l1, l2 := 0, 0
	cur1, cur2 := pHead1, pHead2
	for cur1 != nil {
		l1++
		cur1 = cur1.Next
	}
	for cur2 != nil {
		l2++
		cur2 = cur2.Next
	}

	for i := 0; i < l1+l2 && (pHead1 != nil) && (pHead2 != nil); i++ {
		//相当于把第二个链表的长度加到第一个上面,这样二者就对齐了
		if i > l2 {
			pHead1 = pHead1.Next
		}
		//第一个链表的长度加到第二个上面
		if i > l1 {
			pHead2 = pHead2.Next
		}

		if pHead1 == pHead2 {
			return pHead1
		}
	}
	return nil
}
