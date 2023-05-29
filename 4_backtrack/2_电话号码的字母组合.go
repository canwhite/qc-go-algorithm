package backtrack

var phoneMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

// 入口函数
func LetterCombinations(digits string) []string {

	//定义一个切片，存放最终的结果
	ans := []string{}

	//判断一下边界问题
	if len(digits) == 0 {
		return ans
	}

	//定义一个切片，存放当前的组合
	path := []byte{}

	//调用回调函数，从第0个数字开始遍历
	backtrack(digits, 0, path, &ans)

	return ans

}

// 递归
// 如果加上*可以改外边的值
func backtrack(digits string, index int, path []byte, ans *[]string) {
	//遍历完所有的数字，如果找到了有效的组合，将其加入结果集
	//如果digits是"12",那这个index == 2
	if index == len(digits) {
		*ans = append(*ans, string(path))
		return
	}
	//获取当前数字对应的字母字符串
	letters := phoneMap[digits[index]]

	//遍历每个字母，走位当前位置的选择

	for i := 0; i < len(letters); i++ {
		//将字母加入组合
		path = append(path, letters[i])
		//递归的遍历下一个数字
		backtrack(digits, index+1, path, ans)
		//回溯，撤销选择
		path = path[:len(path)-1]
	}
}
