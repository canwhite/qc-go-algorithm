package ieslide

/*

1. 最大连续子数组和:给定一个整数数组 nums,找出一个连续子数组中有最大和。
使用滑动窗口,当前最大和可以从当前窗口继承,也可以重新开始。
用go实现，代码里加上注释

理解题目：使用滑动窗口,当前最大和可以从当前窗口继承,也可以重新开始的理解
1. 继承前一个窗口的和。如果前一个窗口的和大于 0,我们将它添加到当前窗口中,作为当前窗口的初始和。
2. 重新开始。如果前一个窗口的和小于等于 0,我们将它舍弃,当前窗口的和重新从 0 开始计算。
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxSubArray(nums []int) int {
	// 当前最大和
	maxSum := 0
	// 滑动窗口当前和
	currentSum := 0

	for _, n := range nums {
		// 当前窗口和如果<=0,舍弃当前窗口,从当前元素开始新的窗口
		if currentSum <= 0 {
			currentSum = n
		} else {
			// 否则,当前窗口和加上当前元素
			currentSum += n
		}
		// 更新最大和
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}
