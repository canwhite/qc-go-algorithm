package sort

/**
/ 归并排序的原理:
1. 将数组分成相等的两半
2. 对两半分别进行归并排序
3. 将排序好的两半合并成一个有序的数组
这个过程可以递归的进行，直到数组长度变为1和0
时间复杂度是nlogn
*/

func MergeSort(nums []int) []int {
	//如果切片的长度小于等于1，直接返回，不需要排序
	if len(nums) <= 1 {
		return nums
	}

	//找到切面的中间点，将切片分成左右两部分
	//一般会向下取整，先分再治，这里是分
	mid := len(nums) / 2
	//取左右两边
	left := nums[:mid]
	right := nums[mid:]

	//对左右两边递归的进行归并排序
	//这里是分别对左右的指针先分再治，都是有序的了
	left = MergeSort(left)
	right = MergeSort(right)

	// 最后是合并
	return merge(left, right)

}

// merge函数将两个有序的证书切片合成一个有序的整数切片
func merge(left, right []int) []int {
	//创建一个新的切片，用于存储合并的结果
	result := make([]int, 0, len(left)+len(right))
	//使用指针i和j分别指向左右两个切片的第一个元素
	i := 0
	j := 0
	//当两个指针都没有越界的时候，循环比较它们指向的元素，
	//将较小的元素追加到结果切片中
	//在go里for可以起到while的作用
	for i < len(left) && j < len(right) {
		//然后左右比较去合并
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	//当左边的指针越界时，将右边剩余的值加到结果切片中
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	//当右边的指针越界时，将左边剩余的值加到结果
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	return result
}
