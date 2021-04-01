package main

func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(m, n *TreeNode) bool {
	if m == nil && n == nil {
		return true
	}
	if m == nil {
		return false
	}
	if n == nil {
		return false
	}
	return m.Val == n.Val && isMirror(m.Left, n.Right) && isMirror(m.Right, n.Left)
}

func main() {

}
