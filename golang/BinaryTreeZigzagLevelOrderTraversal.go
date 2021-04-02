package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root != nil {
		u2dII(&res, root, 0)
	}
	return res
}
func u2dII(res *[][]int, node *TreeNode, i int) {
	if node == nil {
		return
	}
	if i >= len(*res) {
		*res = append(*res, make([]int, 0))
	}
	(*res)[i] = append((*res)[i], node.Val)
	if i%2 != 0 {
		copy((*res)[i][1:], (*res)[i])
		(*res)[i][0] = node.Val
	}
	u2dII(res, node.Left, i+1)
	u2dII(res, node.Right, i+1)
}

func main() {

}
