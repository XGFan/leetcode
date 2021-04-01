package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ints := make([]int, 0)
	ints = append(ints, inorderTraversal(root.Left)...)
	ints = append(ints, root.Val)
	ints = append(ints, inorderTraversal(root.Right)...)
	return ints
}

func main() {

}
