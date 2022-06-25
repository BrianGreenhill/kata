package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertbinarytree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = invertbinarytree(root.Right)
	root.Right = invertbinarytree(root.Left)
	return root
}
