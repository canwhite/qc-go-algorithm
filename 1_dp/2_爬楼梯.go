package dp

//爬楼梯: 一次可以爬1级或2级台阶,求n级台阶有多少种爬楼方法。
//解:也使用动态规划,f(n) = f(n-1) + f(n-2)

/**
（1）定义：第i层阶梯有n中爬楼方式
 (2) 递推公式，很明显就是f(n) = f(n-1) + f(n-2)
 (3) 初始化数据，是不是可以初始化到前3
 (4) 确认递推顺序
*/

func ClimbStairs(n int) int {
	//我们来继续实现
	if n <= 3 {
		switch n {
		case 1:
			return 1
		case 2:
			return 2
		case 3:
			return 3
		}
	}
	//初始化前两级台阶的解，用于后续运算
	first := 1
	second := 2

	//使用动态规划
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}
