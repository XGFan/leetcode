package main

func deleteDuplicates_(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for p := head; p.Next != nil; {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}
