package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func nodeToSlice(nodes []*TreeNode, result [][]int) [][]int {
	currentValues := make([]int, len(nodes), len(nodes))
	childrenNodes := make([]*TreeNode, 0, 2*len(nodes))
	reverse := len(result)%2 == 1
	for i, node := range nodes {
		if reverse {
			currentValues[len(nodes)-1-i] = node.Val
		} else {
			currentValues[i] = node.Val
		}
		if node.Left != nil {
			childrenNodes = append(childrenNodes, node.Left)
		}
		if node.Right != nil {
			childrenNodes = append(childrenNodes, node.Right)
		}
	}
	result = append(result, currentValues)
	if len(childrenNodes) == 0 {
		return result
	}
	return nodeToSlice(childrenNodes, result)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	return nodeToSlice([]*TreeNode{root}, make([][]int, 0))
}
