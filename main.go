package main

func main() {
	a := ListNode{5, nil}
	b := ListNode{4, &a}
	c := ListNode{3, &b}
	d := ListNode{2, &c}
	e := ListNode{1, &d}
	reverseKGroup(&e, 2)
}
