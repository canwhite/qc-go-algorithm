package main

import (
	"fmt"
	dp "qc-go-algorithm/1_dp"
	slide "qc-go-algorithm/2_slide"
	tree "qc-go-algorithm/4_btree"
	lib "qc-go-algorithm/lib"
)

func main() {
	// testDp()
	// testSlide()
	testBTree()
}

func testDp() {
	//1.斐波那契数
	result := dp.Fibonacci(9)
	fmt.Println(result)

}
func testSlide() {
	//最大连续子数组和
	nums := []int{1, 2, -3, -4, 3, 0, 7, 8, 2, 10}
	max := slide.MaxSubArray(nums)
	fmt.Println(max)

}

func testBTree() {
	//需要一个root值

	t1 := &lib.TreeNode{Val: 1}

	t2 := &lib.TreeNode{Val: 2}
	t1.Left = t2

	t3 := &lib.TreeNode{Val: 3}
	t1.Right = t3

	t4 := &lib.TreeNode{Val: 4}
	t2.Left = t4

	t5 := &lib.TreeNode{Val: 5}
	t3.Right = t5

	t6 := &lib.TreeNode{Val: 6}
	t4.Left = t6

	t7 := &lib.TreeNode{Val: 7}
	t5.Right = t7

	result := tree.MaxDepth(t1)
	fmt.Println(result)

}
