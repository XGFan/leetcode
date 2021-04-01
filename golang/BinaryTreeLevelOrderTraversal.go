package main

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root != nil {
		u2d(&res, []*TreeNode{root})
	}
	return res
}
func u2d(res *[][]int, nodes []*TreeNode) {
	if len(nodes) == 0 {
		return
	}
	treeNodes := make([]*TreeNode, 0)
	ints := make([]int, 0)
	for _, v := range nodes {
		ints = append(ints, v.Val)
		if v.Left != nil {
			treeNodes = append(treeNodes, v.Left)
		}
		if v.Right != nil {
			treeNodes = append(treeNodes, v.Right)
		}
	}
	*res = append(*res, ints)
	u2d(res, treeNodes)
}

func main() {

}
