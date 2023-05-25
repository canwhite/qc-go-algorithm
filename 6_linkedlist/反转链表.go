package linkedlist

//除了这种直接倒入包的方法，还可以像python一样，直接导入方法
import (
	lib "qc-go-algorithm/lib"
)

// 代码详解 ， 看这里，https://juejin.cn/post/6844904058562543623
func ReverseList(head *lib.ListNode) *lib.ListNode {

	//先习惯性判断边界情况
	//如果链表为空，或者只有一个节点，直接返回
	if head == nil || head.Next == nil {
		return head
	}
	//定义prev和curr
	var prev *lib.ListNode = nil
	curr := head

	//因为curr指向了head，head不为nil，所以进入循环
	//eg ： 1 -> 2 -> 3 -> 4 -> 5
	for curr != nil {
		//循环一：第一次nextTemp = 2
		//循环二；nextTemp = 3； curr指向2，prev指向1，在另一个链表里
		nextTemp := curr.Next //现表指针

		//循环一：把1拔出来，让它的next指向nil，单独成一条链
		//循环二：将2把出来，放在另一个链表里1的前边
		curr.Next = prev //拔头成新表

		//循环一：然后把curr赋值给prev，这样cur和prev都指向1
		//循环二：让cur指向单独链表的头部
		prev = curr //将prev放在新表头

		//循环一：把nextTemp赋值给cur，即curr指向nextTemp，然后prev在1的位置
		//循环二：将curr指向3
		curr = nextTemp //curr移动到现表，继续下个循环拔头
	}

	//总结：就是不断利用头插法，遍历每个元素，不断在新链表头部插入，实现翻转

	return prev
}
