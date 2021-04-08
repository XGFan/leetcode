package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	return pathSumIter([]int{}, root, targetSum)
}

func pathSumIter(prefix []int, root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	nPrefix := make([]int, len(prefix)+1)
	copy(nPrefix, prefix)
	nPrefix[len(prefix)] = root.Val
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return [][]int{nPrefix}
		} else {
			return nil
		}
	} else {
		wanted := targetSum - root.Val
		if root.Left != nil && root.Right == nil {
			return pathSumIter(nPrefix, root.Left, wanted)
		}
		if root.Right != nil && root.Left == nil {
			return pathSumIter(nPrefix, root.Right, wanted)
		}
		return append(pathSumIter(nPrefix, root.Left, wanted), pathSumIter(nPrefix, root.Right, wanted)...)
	}
}

func main() {

}
