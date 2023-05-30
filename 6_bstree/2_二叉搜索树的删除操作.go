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

// 定义一个函数，用于从二叉搜索树中删除一个节点
// key是要删除的值
func DeleteNode(root *lib.TreeNode, key int) *lib.TreeNode {
	//还是先判断边界问题
	if root == nil {
		return nil
	}

	//如果要删除的节点值小于根节点值，说明在左子树中，
	//递归的在左子树中删除，并更新左子树指针
	if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
	}
	//如果要删除的节点大于根节点值，说明在右子树中，
	//递归的在右子树中删除，并更新右子树指针
	if key > root.Val {
		root.Right = DeleteNode(root.Right, key)
	}

	//如果要删除的节点等于根节点，说明找到了目标节点，分三种情况讨论
	if key == root.Val {
		//1.如果目标节点没有左右子树，直接删除它，返回nil
		if root.Left == nil && root.Right == nil {
			root = nil
			return root
		}

		//2.如果目标节点只有右子树，用右子树替换它，返回右子树指针
		if root.Right == nil {
			root = root.Right
			return root
		}
		//3.如果目标节点只有左子树。用左子树替换它，返回左子树指针
		if root.Left == nil {
			root = root.Left
			return root
		}

		//4.如果目标节点有左右子树，用右子树的最小节点替换它，并删除右子树的最小节点，返回更新后的根节点指针
		minNode := findMin(root.Right)
		root.Val = minNode.Val
		//赋值完，删除右子树的最小节点
		root.Right = DeleteNode(root.Right, minNode.Val)
	}

	return root
}
