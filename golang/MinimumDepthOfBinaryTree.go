package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		l := minDepth(root.Left)
		r := minDepth(root.Right)
		if l < r {
			return l + 1
		} else {
			return r + 1
		}
	} else {
		if root.Left != nil {
			return minDepth(root.Left) + 1
		} else {
			return minDepth(root.Right) + 1
		}
	}
}

func main() {

}
