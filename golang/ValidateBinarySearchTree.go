package main

import "fmt"

func isValidBST(root *TreeNode) bool {
	return isValidBSTIter(-1<<31-1, 1<<31, root)
}

func isValidBSTIter(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return isValidBSTIter(min, root.Val, root.Left) && isValidBSTIter(root.Val, max, root.Right)
}

func main() {
	fmt.Println(1<<30 - 1 + 1<<30)
}
