package main

func main() {
	a, b, c, d, e := ListNode{1, nil},
		ListNode{2, nil},
		ListNode{3, nil},
		ListNode{4, nil},
		ListNode{5, nil}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e

	reverseKGroup(&a, 2)
}
