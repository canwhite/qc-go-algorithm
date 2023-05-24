package bstree

import (
	lib "qc-go-algorithm/lib"
)

func InsertIntoBST(root *lib.TreeNode, val int) *lib.TreeNode {

	//如果树是空的，直接插入根节点
	if root == nil {
		return &lib.TreeNode{Val: val}
	}

	return nil
}
