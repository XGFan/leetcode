package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func recoverTree(root *TreeNode) {
	i := -1 << 32
	var ints []int
	findTrouble(&ints, &i, root)
	m, n := ints[0], ints[0]
	for j := 1; j < len(ints); j++ {
		if m < ints[j] {
			m = ints[j]
		}
		if n > ints[j] {
			n = ints[j]
		}
	}
	//fmt.Println(m, n)
	simpleRecover(root, m, n)
}

func simpleRecover(root *TreeNode, m, n int) {
	if root != nil {
		if root.Val == m {
			root.Val = n
		} else if root.Val == n {
			root.Val = m
		}
		simpleRecover(root.Left, m, n)
		simpleRecover(root.Right, m, n)
	}
}

func findTrouble(res *[]int, i *int, root *TreeNode) {
	if root != nil {
		findTrouble(res, i, root.Left)
		if *i > root.Val {
			*res = append(*res, root.Val, *i)
		}
		*i = root.Val
		findTrouble(res, i, root.Right)
	}
}

func main() {
	recoverTree(&TreeNode{Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}, Val: 1})
	fmt.Println()
	recoverTree(&TreeNode{Left: &TreeNode{Val: 1}, Val: 3, Right: &TreeNode{Left: &TreeNode{Val: 2}, Val: 4}})
	fmt.Println()
	recoverTree(&TreeNode{
		Left: &TreeNode{
			Left:  &TreeNode{Val: 1},
			Val:   2,
			Right: &TreeNode{Val: 6},
		},
		Val: 4,
		Right: &TreeNode{
			Left:  &TreeNode{Val: 5},
			Val:   3,
			Right: &TreeNode{Val: 7},
		},
	})
	fmt.Println()
	recoverTree(&TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Left:  &TreeNode{Val: 1},
				Val:   2,
				Right: &TreeNode{Val: 6},
			},
			Val: 4,
			Right: &TreeNode{
				Left:  &TreeNode{Val: 5},
				Val:   3,
				Right: &TreeNode{Val: 7},
			},
		},
		Val: 8,
		Right: &TreeNode{
			Left: &TreeNode{
				Left:  &TreeNode{Val: 9},
				Val:   10,
				Right: &TreeNode{Val: 11},
			},
			Val: 12,
			Right: &TreeNode{
				Left:  &TreeNode{Val: 13},
				Val:   14,
				Right: &TreeNode{Val: 15},
			},
		},
	})
}
