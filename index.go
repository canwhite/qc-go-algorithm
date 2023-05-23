package main

import (
	"fmt"
	dp "qc-go-algorithm/1_dp"
)

func main() {
	testDp()
}

func testDp() {
	//1.斐波那契数
	result := dp.Fibonacci(9)
	fmt.Println(result)
}
