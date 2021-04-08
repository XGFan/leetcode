package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		t := root.Left
		for t.Right != nil {
			t = t.Right
		}
		t.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}

func main() {

}
