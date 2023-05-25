package backtrack

/**
遍历子集的问题
回溯算法的一般流程是，for循环横向遍历，递归纵向遍历。用于解决一些组合切割子集的问题
void backtracking(参数) {
	if (终止条件) {
		存放结果;
		return;
	}

	for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
		处理节点;
		backtracking(路径，选择列表); // 递归
		回溯，撤销处理结果
	}
}
*/
