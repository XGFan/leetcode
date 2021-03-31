package main

func partition(head *ListNode, x int) *ListNode {
	tail := &ListNode{}
	prev := &ListNode{Next: tail}
	res := prev
	for head != nil {
		if head.Val >= x {
			tail.Next, head, head.Next, tail = head, head.Next, tail.Next, head
		} else {
			prev.Next, head, head.Next, prev = head, head.Next, prev.Next, head
		}
	}
	prev.Next = prev.Next.Next
	return res.Next
}

func main() {

}
