package main

func generateTrees(n int) []*TreeNode {
	ints := make([]int, n)
	for i := range ints {
		ints[i] = i + 1
	}
	return generateTreesIter(ints)
}

func generateTreesIter(ints []int) []*TreeNode {
	if len(ints) == 0 {
		return []*TreeNode{nil}
	}
	res := make([]*TreeNode, 0)
	for i, v := range ints {
		ltrees := generateTreesIter(ints[:i])
		rtrees := generateTreesIter(ints[i+1:])
		for _, l := range ltrees {
			for _, r := range rtrees {
				res = append(res, &TreeNode{
					Left:  l,
					Val:   v,
					Right: r,
				})
			}
		}
	}
	return res
}

func main() {

}
