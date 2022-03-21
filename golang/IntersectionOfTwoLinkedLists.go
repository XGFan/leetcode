package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p := headA
	for p != nil {
		p.Next = &ListNode{Next: p.Next}
		p = p.Next.Next
	}
	var result *ListNode
	q := headB
	for q != nil {
		if q.Next != nil && q.Next.Val == 0 {
			result = q
			break
		}
		q = q.Next
	}
	p = headA
	for p != nil {
		if p.Next != nil {
			p.Next = p.Next.Next
		}
		p = p.Next
	}
	return result
}

func main() {

}
