package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	n := 0
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		barrier := len(nodes)
		for i := 0; i < barrier; i++ {
			node := nodes[i]
			if node.Left == nil && node.Right == nil {
				return n + 1
			} else if node.Left != nil && node.Right == nil {
				nodes[i] = node.Left
			} else if node.Right != nil && node.Left == nil {
				nodes[i] = node.Right
			} else {
				nodes[i] = node.Left
				nodes = append(nodes, node.Right)
			}
		}
		n++
	}
	return n
}

func main() {

}
