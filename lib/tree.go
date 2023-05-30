package lib

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 定义一个辅助函数用于打印tree
func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	//然后递归输出
	PrintTree(root.Left)
	fmt.Print(root.Val, " ")
	PrintTree(root.Right)
}
