package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	} else {
		wanted := targetSum - root.Val
		if root.Left != nil && root.Right == nil {
			return hasPathSum(root.Left, wanted)
		}
		if root.Right != nil && root.Left == nil {
			return hasPathSum(root.Right, wanted)
		}
		return hasPathSum(root.Left, wanted) || hasPathSum(root.Right, wanted)
	}
}

func main() {

}
