package linkedlist

import lib "qc-go-algorithm/lib"

// 判断一个链表是否是回文链表的函数
func IsPalindrome(head *lib.ListNode) bool {

	//如果链表为空或者只有一个节点，直接返回为true
	if head == nil || head.Next == nil {
		return true
	}

	// 使用快慢指针找到链表的中点
	slow, fast := head, head

	// 当快指针没有到达尾节点时，循环继续
	// 因为快指针每次走两步，慢指针每次走一步，所以当快指针到尾巴的时候，
	// 慢指针刚好走了一半的距离，也就是到了中点。这是一种常用的找链表中点的方法。
	for fast.Next != nil && fast.Next.Next != nil {
		//快指针每次走一步
		slow = slow.Next
		//快指针每次走两步
		fast = fast.Next.Next
	}
	//反转后半部分链表，从中点的下一个节点开始
	//定义一个前驱节点，初始为nil
	var prev *lib.ListNode
	//定义一个当前节点，初始为中点的下一个节点
	curr := slow.Next

	//当当前节点不为空时，循环继续
	for curr != nil {
		//这部分相当于头插法，将cur移动到了nil的前边，然后将后边的依次往另外一条链的前边放
		//就实现了反转代码的作用
		next := curr.Next
		//相当于curr的next变成了prev，相当于单独走一条链
		//将curr放在了nil的前边
		curr.Next = prev
		//更新前驱节点到新起链表的头部
		prev = curr
		//然后将cur指向原来链表的next，开始新的一次循环
		curr = next
	}

	//比较前半段和后半段链表的节点值，如果不同，返回为false
	//定义两个指针，分别指向前半段和后半段的头节点
	p1, p2 := head, prev

	//当两个指针不为空时，循环继续
	for p1 != nil && p2 != nil {
		//如果两个指针指向的节点值不相同，说明不是回文链表，返回false
		if p1.Val != p2.Val {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next
	}
	//如果没有不同，返回为true
	return true
}
