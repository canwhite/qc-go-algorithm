package sort

/*
*
堆排序的原理是利用一个最大堆或者最小堆来对一个数组进行排序。
一个最大堆是一个完全二叉树，满足每个节点的值都大于等于它的子节点的值。
一个最小堆是一个完全二叉树，满足每个节点的值都小于等于它的子节点的值。

PS-1：完全二叉树和满二叉树的区别
完全二叉树
 1. 除了最后一层,其他层的节点数达到最大值。
 2. 最后一层或者连续多层的叶节点在左边连续出现。
    1
    /   \
    2     3
    / \   /

4   5 6

满二叉树:
 1. 所有层的节点数都达到最大值。
 2. 叶节点只能出现在最后两层。
    1
    /     \
    2       3
    / \     / \
    4   5   6   7

PS-2：平衡二叉树是一种完全二叉树吗？

平衡二叉树不一定是完全二叉树，因为它可能有空缺的节点，或者最后一层的节点不是靠左排列的

	   1
	  / \
	 2   3
	/ \   \

4   5   6

---------------------------------算法公式--------------------------------------------

堆排序的过程分为两个步骤：

建堆：将一个无序的数组调整成一个最大堆或者最小堆，这个过程可以自底向上或者自顶向下进行，时间复杂度是O(n)。

排序：每次从堆中取出堆顶元素（最大值或者最小值），放到数组的末尾，然后将剩余的元素重新调整成一个堆，

	这个过程需要重复n-1次，每次调整的时间复杂度是O(logn)，所以总的时间复杂度是O(nlogn)。

这是因为完全二叉树的每个节点都有两个子节点（除了叶子节点），
所以它的子节点的序号是
如果父节点是i，
i*2+1  或者  i*2+2。
你可以用数学归纳法来证明这个公式，或者用自然语言来解释这个公式。

如果子节点是n，
(n-1)/2或者(n-2)/2

如果针对两个，那么我们就要用两个不同的公式，
一个是(n-1)/2，另一个是(n-2)/2。这样就不方便了。我们要的是通用
而且，如果我们用(n-1)/2这个公式，那么无论是左孩子还是右孩子，都可以得到正确的父节点的序号。

例如，如果一个节点的序号是4，那么它的父节点的序号就是(4-1)/2=1.5，取整数部分就是1。
如果一个节点的序号是5，那么它的父节点的序号也是(5-1)/2=2，取整数部分也是2。
所以我们可以用一个公式来适用于所有的子节点。

	   A(0)
	  /   \
	B(1)  C(2)
	/ \   / \

D(3)E(4)F(5)G(6)

*****
*/

/*
---------------------------------具体操作--------------------------------------------
上下两个for循环有哪些区别？

上面的for循环是用来建立最大堆的：
它从最后一个非叶子节点开始，自底向上地调整每个节点，使其满足最大堆的性质。

-----ps-1 为什么要从最后一个非叶子节点开始
第一个for循环要从最后一个非叶子节点开始，是因为叶子节点本身就满足最大堆的性质，不需要调整。
而非叶子节点可能需要调整，使其大于或等于它的子节点。
从最后一个非叶子节点开始，自底向上地调整，可以保证每次调整后，它的子树都是一个最大堆。

-----ps-2 那为什么叶子节点本身就满足最大堆的性质呢？

叶子节点本身就满足最大堆的性质，是因为它们没有子节点，所以不需要和子节点比较大小。
最大堆的性质是每个节点都大于或等于它的子节点，而叶子节点没有子节点，所以它们自然就满足了这个性质。

下面的for循环是用来排序的：
它每次将堆顶元素（最大值）与堆的最后一个元素交换，并从堆中移除，然后对剩余的元素重新调整成最大堆。
这样就可以保证每次交换后，堆的最后一个元素是当前最大值。

-----ps-1 第二个for循环是怎么从堆中移除元素的呢？

第二个for循环是这样从堆中移除元素的：

首先，将堆顶元素（最大值）和堆的最后一个元素交换，这样就相当于将最大值移动到数组的末尾，并从堆中移除。
然后，将len参数减一，表示堆的大小减少了一个元素，也就是忽略了数组的最后一个元素。
最后，对堆顶元素进行调整，使其满足最大堆的性质，这样就完成了一次移除操作。
重复这个过程，直到堆为空，就可以得到排序后的数组。

--ps-1
*/
func HeapSort(nums []int) []int {
	// 建立初始堆,i表示第一个非叶子节点,len(nums)/2-1表示最后一个非叶子节点
	//这里因为i是从0开始的，所以（n-1-1）/2 就能与n/2-1
	for i := len(nums)/2 - 1; i >= 0; i-- {
		// 调整堆,将最大值移至顶部
		heapify(nums, i, len(nums))
	}
	// 移除堆顶元素并调整堆,重复直到堆为空
	// 移除栈顶元素的方法，就是当前长减去1
	// 这里最初只是一个初始值，这个i十一直变化的
	// 通过交换和减少len参数来实现从堆中移除元素的效果
	for i := len(nums) - 1; i > 0; i-- {
		// 先交换，再减少len
		nums[0], nums[i] = nums[i], nums[0]
		// 调整堆。len参数减去1发生在这里
		heapify(nums, 0, i)
	}
	return nums
}

func heapify(nums []int, i, len int) {
	// 取出i节点的左右子节点
	l := 2*i + 1
	r := 2*i + 2
	// 初始化最大值索引为i
	largest := i
	// 如果左子节点大于i且存在左子节点,则将largest设置为左子节点
	if l < len && nums[l] > nums[largest] {
		largest = l
	}
	// 如果右子节点大于largest索引指向的节点,将largest设置为右子节点
	if r < len && nums[r] > nums[largest] {
		largest = r
	}
	// 如果存在更大的子节点,则与i节点交换并继续调整堆
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, largest, len)
	}
}
