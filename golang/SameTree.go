package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		if q == nil {
			return true
		}
		return false
	} else {
		if q == nil {
			return false
		} else {
			return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}
}

func main() {

}
