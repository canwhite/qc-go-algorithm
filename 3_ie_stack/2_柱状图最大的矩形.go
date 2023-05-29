package iestack

/**
  先看一种理解
  看了别人的答案想了半天才明白……其实可以把这个想象成锯木板，如果木板都是递增的那我很开心，
  如果突然遇到一块木板i矮了一截，那我就先找之前最戳出来的一块（其实就是第i-1块），
  计算一下这个木板单独的面积，然后把它锯成次高的，
  这是因为我之后的计算都再也用不着这块木板本身的高度了。
  再然后如果发觉次高的仍然比现在这个i木板高，
  那我继续单独计算这个次高木板的面积（应该是第i-1和i-2块），再把它俩锯短。
  直到发觉不需要锯就比第i块矮了，那我继续开开心心往右找更高的。
  当然为了避免到了最后一直都是递增的，所以可以在最后加一块高度为0的木板。
  这个算法的关键点是把那些戳出来的木板早点单独拎出来计算，
  然后就用不着这个值了。说实话真的很佩服第一个想出来的人……
*/

func LargestRectangleArea(heights []int) int {
	// 预置一个结果集
	area := 0
	// 在柱体的数组的头和尾加了两个高度为0的柱体
	tmp := make([]int, len(heights)+2)
	// copy heights的值
	copy(tmp[1:], heights)
	// 定义一个栈，存放柱子的索引
	stack := []int{}
	// 遍历每个柱子

	for i := 0; i < len(tmp); i++ {
		num := tmp[i]
		// 如果num比栈顶元素小，就弹出栈顶元素，让小数下沉，就得到一个单调递减栈
		// 当前边的值大于当前值的时候，去计算，前边找大的
		// 栈不为空，且后边出现了小值
		for len(stack) > 0 && num < tmp[stack[len(stack)-1]] {
			// 弹出栈顶元素
			x := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 计算以栈顶元素为高的矩形的面积，并更新结果集
			h := tmp[x] //拿到实际存储的前边的值
			//计算面积，两点之间的距离是差-1
			area = max(area, (i-stack[len(stack)-1]-1)*h)
		}
		// i入栈，看着存的是下标，实际上存的是值啊
		stack = append(stack, i)
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
