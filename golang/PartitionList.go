package main

func partition(head *ListNode, x int) *ListNode {
	tail := &ListNode{}
	p := &ListNode{Next: tail}
	res := p
	for head != nil {
		if head.Val >= x {
			tail.Next = head
			head = head.Next
			tail = tail.Next
			tail.Next = nil
		} else {
			t := head
			head = head.Next
			t.Next = p.Next
			p.Next = t
			p = p.Next
		}
	}
	if p.Next != nil {
		p.Next = p.Next.Next
	}
	return res.Next
}

func main() {

}
