package main

/*
 给一个长度为n链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null。
*/
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	store := make(map[*ListNode]struct{})
	for pHead != nil {
		if _, exist := store[pHead]; exist {
			return pHead
		} else {
			store[pHead] = struct{}{}
		}
		pHead = pHead.Next
	}

	return nil
}

// func EntryNodeOfLoop(pHead *ListNode) *ListNode {
// 	if pHead == nil {
// 		return nil
// 	}
// 	fast, slow := pHead, pHead
// 	for fast != nil && slow != nil {
// 		slow = slow.Next
// 		fast = fast.Next
// 		if fast != nil {
// 			fast = fast.Next
// 		}
// 		if fast == slow {
// 			break
// 		}
// 	}
// 	if fast == nil || slow == nil {
// 		return nil
// 	}

// 	fast = pHead
// 	for fast != nil && slow != nil {
// 		fast = fast.Next
// 		slow = slow.Next
// 		if fast == slow {
// 			break
// 		}
// 	}

// 	return fast
// }
