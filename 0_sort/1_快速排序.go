package sort

//快速排序，因为间接是个树形，所以时间复杂度整体是nlogn

// 定义一个快排函数，接收一个切片作为参数，返回一个排序后的切片
func QuickSort(nums []int) []int {
	//如果切片的长度小于1，直接返回，不需要排序
	if len(nums) < 1 {
		return nums
	}
	//定义一个基准值的索引，取切片的中间位置
	pivotIndex := len(nums) / 2

	//定义一个基准值，取切片中间位置的元素，并将其从切片中删除
	pivot := nums[pivotIndex]
	//并将其从切片中删除
	//...表示保存原有容量，不用新建，以提高性能
	nums = append(nums[:pivotIndex], nums[pivotIndex+1:]...)

	//定义两个切片，用于存放左边的和右边的元素
	var left []int
	var right []int

	//遍历切片，从左到右
	for i := 0; i < len(nums); i++ {
		//如果当前元素小于基准值
		if nums[i] < pivot {
			//则将当前元素追加到左边的切片中
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}

	}

	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}
