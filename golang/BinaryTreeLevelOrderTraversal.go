package main

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root != nil {
		u2d(&res, []*TreeNode{root})
	}
	return res
}
func u2d(res *[][]int, nodes []*TreeNode) {
	size := len(nodes)
	if size == 0 {
		return
	}
	ints := make([]int, 0, size)
	for i := 0; i < size; i++ {
		v := nodes[i]
		ints = append(ints, v.Val)
		if v.Left != nil {
			nodes = append(nodes, v.Left)
		}
		if v.Right != nil {
			nodes = append(nodes, v.Right)
		}
	}
	*res = append(*res, ints)
	u2d(res, nodes[:copy(nodes, nodes[size:])])
}

func main() {

}
