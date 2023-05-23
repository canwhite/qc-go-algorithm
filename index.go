package main

import (
	"fmt"
	dp "qc-go-algorithm/1_dp"
	slide "qc-go-algorithm/2_slide"
)

func main() {
	// testDp()
	testSlide()
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
