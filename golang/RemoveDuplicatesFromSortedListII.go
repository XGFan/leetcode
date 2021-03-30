package main

func deleteDuplicates(head *ListNode) *ListNode {
	duplicated := false
	p := &ListNode{Next: head}
	var prev = p
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			duplicated = true
			head.Next = head.Next.Next //删除next节点
		} else {
			if duplicated {
				prev.Next = head.Next //删除当前节点
			} else {
				prev = head
			}
			head = head.Next
			duplicated = false
		}
	}
	return p.Next
}

func main() {

}
