package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if r < r {
		return r + 1
	} else {
		return l + 1
	}
}

func main() {

}
