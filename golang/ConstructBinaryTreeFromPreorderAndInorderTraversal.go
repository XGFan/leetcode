package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := preorder[0]
	i := 0
	for i = range inorder {
		if inorder[i] == root {
			break
		}
	}
	return &TreeNode{
		Val:   root,
		Left:  buildTree(preorder[1:i+1], inorder[:i]),
		Right: buildTree(preorder[i+1:], inorder[i+1:]),
	}
}

func main() {

}
