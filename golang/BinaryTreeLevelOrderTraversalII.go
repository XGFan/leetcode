package main

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root != nil {
		u2d2(&res, &[]*TreeNode{root})
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func u2d2(res *[][]int, nodes *[]*TreeNode) {
	size := len(*nodes)
	if size == 0 {
		return
	}
	ints := make([]int, 0, size)
	for i := 0; i < size; i++ {
		v := (*nodes)[i]
		ints = append(ints, v.Val)
		if v.Left != nil {
			*nodes = append(*nodes, v.Left)
		}
		if v.Right != nil {
			*nodes = append(*nodes, v.Right)
		}
	}
	*res = append(*res, ints)
	*nodes = (*nodes)[:copy(*nodes, (*nodes)[size:])]
	u2d2(res, nodes)
}

func main() {

}
