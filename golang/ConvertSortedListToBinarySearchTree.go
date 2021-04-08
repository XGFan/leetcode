package main

func sortedListToBST(head *ListNode) *TreeNode {
	i := 0
	t := head
	for t != nil {
		t = t.Next
		i++
	}
	tree, _ := buildBST(head, i)
	return tree
}

func buildBST(head *ListNode, n int) (*TreeNode, *ListNode) {
	if n == 0 {
		return nil, nil
	}
	if n == 1 {
		return &TreeNode{
			Val: head.Val,
		}, head.Next
	}
	left, listNode := buildBST(head, n/2)
	right, tail := buildBST(listNode.Next, n-1-n/2)
	return &TreeNode{
		Val:   listNode.Val,
		Left:  left,
		Right: right,
	}, tail
}

func main() {

}
