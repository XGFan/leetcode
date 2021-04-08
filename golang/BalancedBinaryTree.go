package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isBalanced(root *TreeNode) bool {
	t, _ := balanceAndTall(root)
	return t
}

func balanceAndTall(root *TreeNode) (c bool, tall int) {
	if root == nil {
		return true, 0
	}
	lb, ltall := balanceAndTall(root.Left)
	if !lb {
		return lb, 0
	}
	rb, rtall := balanceAndTall(root.Right)
	if !rb {
		return rb, 0
	}

	if ltall+1 < rtall || rtall+1 < ltall {
		return false, 0
	} else {
		if ltall > rtall {
			return true, ltall + 1
		} else {
			return true, rtall + 1
		}
	}
}

func main() {
	fmt.Println(isBalanced(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}))

}
