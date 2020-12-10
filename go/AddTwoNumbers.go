package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans = l1
	var f = 0
	for {
		sum := l1.Val + l2.Val + f
		l1.Val = sum % 10
		f = sum / 10
		if l2.Next == nil {
			if f == 0 {
				return ans
			}
			l2.Next = &ListNode{}
		}
		if l1.Next == nil {
			l1.Next = l2.Next
			if f == 0 {
				return ans
			}
			l2.Next = &ListNode{}
		}
		l1 = l1.Next
		l2 = l2.Next
	}
}
