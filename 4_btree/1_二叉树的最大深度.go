package btree

//PS：这里引入和外层统一，因为外层是mod所在的位置
//以此为支点才能更好的寻找
import (
	lib "qc-go-algorithm/lib"
)

func MaxDepth(root *lib.TreeNode) int {
	//递推终止条件
	if root == nil {
		return 0
	}
	//DFS搜索左右子树的最大深度
	leftMax := MaxDepth(root.Left)
	rightMax := MaxDepth(root.Right)
	//取左右子树最大深度的最大值，然后加上当前根节点，得到此树最大深度
	max := leftMax
	if rightMax > leftMax {
		max = rightMax
	}
	return max + 1
}
