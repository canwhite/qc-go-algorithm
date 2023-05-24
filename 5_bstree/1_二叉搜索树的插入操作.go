package bstree

import (
	lib "qc-go-algorithm/lib"
)

func InsertIntoBST(root *lib.TreeNode, val int) *lib.TreeNode {

	//1.如果树是空的，直接插入根节点
	if root == nil {
		//相当于新建一个返回
		return &lib.TreeNode{Val: val}
	}

	//2.如果插入值<=当前节点值，插入到左子树
	if val <= root.Val {
		//2.1如果左树为空，插入左树，如果左树非空，递归插入
		if root.Left == nil {
			root.Left = &lib.TreeNode{Val: val}
		} else {
			root.Left = InsertIntoBST(root.Left, val)
		}
	}
	//3.如果插入值>当前节点值，插入到右子树
	//如果右树为空，插入右树，如果右树非空，递归插入

	if val > root.Val {
		//3.1 如果右树为空，插入右树，如果右树为空，递归插入
		if root.Right == nil {
			root.Right = &lib.TreeNode{Val: val}
		} else {
			root.Right = InsertIntoBST(root.Right, val)
		}
	}
	return root
}
