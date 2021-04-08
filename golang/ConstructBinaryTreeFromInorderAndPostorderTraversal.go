package main

//inorder 左中右
//postorder 左右中
func buildTree_(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := postorder[len(postorder)-1]
	i := 0
	for i = range inorder {
		if inorder[i] == root {
			break
		}
	}
	return &TreeNode{
		Val:   root,
		Left:  buildTree_(inorder[:i], postorder[:i]),
		Right: buildTree_(inorder[i+1:], postorder[i:len(postorder)-1]),
	}
}

func main() {

}
