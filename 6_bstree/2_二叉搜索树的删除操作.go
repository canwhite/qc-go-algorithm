package bstree

import "qc-go-algorithm/lib"

//定义一个辅助函数，用于找到一棵树中的最小节点
//用于找到右子树的最小节点

func findMin(node *lib.TreeNode) *lib.TreeNode {
	//先考虑边界问题
	//如果节点为空，直接返回nil
	if node == nil {
		return nil
	}

	//如果节点没有左子树，说明它就是最小的节点，返回它
	if node.Left == nil {
		return node
	}

	//否则，递归的在左子树中寻找最小节点
	return findMin(node.Left)
}
