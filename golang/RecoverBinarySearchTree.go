package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func recoverTree(root *TreeNode) {
	i := -1 << 32
	ints := []int{1 << 31, -1<<31 - 1}
	treeToSlice(&ints, &i, root)
	recover(root, ints[0], ints[1])
	//fmt.Println(ints)
}

func recover(root *TreeNode, m, n int) {
	if root != nil {
		if root.Val == m {
			root.Val = n
		} else if root.Val == n {
			root.Val = m
		}
		recover(root.Left, m, n)
		recover(root.Right, m, n)
	}
}

func treeToSlice(res *[]int, i *int, root *TreeNode) {
	if root != nil {
		treeToSlice(res, i, root.Left)
		if *i > root.Val {
			if (*res)[0] > root.Val {
				(*res)[0] = root.Val
			}
			if (*res)[1] < *i {
				(*res)[1] = *i
			}
			//*res = append(*res, *i, root.Val)
			//fmt.Println(*i, root.Val)
		}
		*i = root.Val
		treeToSlice(res, i, root.Right)
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
