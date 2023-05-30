package btree

import "qc-go-algorithm/lib"

// 2. 二叉树的最小深度:给定一个二叉树,找出其最小深度。使用DFS或BFS遍历树,维持一个最小深度变量。
// 最小深度是指从根节点到最近的叶子节点的路径上的节点个数。叶子节点是指没有子节点的节点
func MinDepth(root *lib.TreeNode) int {

	//先考虑边界情况
	//如果根节点为空，返回0
	if root == nil {
		return 0
	}

	//定义一个变量ans，记录最小深度，初始化为最大整数
	// 1<< 31 ,将1左移31位，再减去1，就是最大整数
	ans := 1<<31 - 1

	//定义一个辅助函数，用于递归的遍历树，参数是当前节点和当前深度

	var dfs func(node *lib.TreeNode, depth int)
	//然后初始化
	dfs = func(node *lib.TreeNode, depth int) {
		//还是先考虑边界问题
		if node == nil {
			return
		}
		//如果当前节点是叶子节点，说明我们找到了一条从根到叶的路径，更新最小深度
		if node.Left == nil || node.Right == nil {
			if depth < ans {
				ans = depth
			}
			return
		}
		//如果当前节点不是叶子节点，递归的遍历左右子树，深度+1
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	//从根节点开始遍历，初始深度为1
	dfs(root, 1)

	return ans
}
