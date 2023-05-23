package backtrack

/**
1. 组合总和:给定一个无重复元素的数组 candidates 和一个目标数 target ,
找出 candidates 中所有可以使数字和为 target 的组合。
使用回溯法遍历所有可能的组合。用go来实现，注释写进代码里
*/

func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	//如果传入的是result，相当于放一个副本进去
	//&result，改的就是外部这个result，这样就很方便
	backtrace(candidates, target, 0, []int{}, &result)
	return result //都可以先以返回nil取代
}

// path 是满足情况的一个子集
// result *[][]int 这里*[][]int 只是声明是一个指针类型，虽然传入的是地址
func backtrace(candidates []int, target int, index int, path []int, result *[][]int) {

	//如果找到一个解，就往结果集里加，
	if target == 0 {
		//声明一个切片，预置空间
		combi := make([]int, len(path))
		copy(combi, path)
		//对值进行操作
		*result = append(*result, combi)
		return
	}

	//超出范围，直接返回
	if target < 0 {
		return
	}

	//还没有搜索完，继续搜索
	for i := index; i < len(candidates); i++ {
		//做选择
		path = append(path, candidates[i])
		//进入下一层搜索
		backtrace(candidates, target-candidates[i], i, path, result)

		//撤销选择
		path = path[:len(path)-1]
	}

}
