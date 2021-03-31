package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	res := &ListNode{Next: head}
	head = res
	p := res
	for i := 0; i < left; i++ {
		p = head
		head = head.Next
	}
	p.Next = nil
	tail := p.Next
	for i := left; i <= right; i++ {
		p.Next = head
		head = head.Next
		p.Next.Next = tail
		tail = p.Next
	}
	if tail != nil {
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = head
	}
	return res.Next
}

func main() {

}
