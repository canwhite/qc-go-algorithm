package dp

//1. 斐波那契数列: 求斐波那契数列第n项的值。
//解:使用动态规划,有递推公式f(n) = f(n-1) + f(n-2). 用go实现每一步加上注释

func Fibonacci(n int) int {
	//先加个截止
	if n < 0 {
		return -1
	}
	//初始值
	if n == 0 || n == 1 {
		return n
	}
	//递推公式都给到了
	return Fibonacci(n-1) + Fibonacci(n-2)
}
