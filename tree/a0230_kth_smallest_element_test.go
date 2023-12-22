package tree

/*
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。



解题方法:
什么是二叉搜索树：左子树节点 < 根节点 < 右子树节点
二叉搜索树的前序遍历：左子树 -> 根节点 -> 右子树，再结合二叉搜索的性质，刚好是从小到大遍历整棵树
*/

func kthSmallest(root *TreeNode, k int) int {
	node := &TreeNode{}
	cur := 0
	traverse(root, &cur, k, node)
	if node != nil {
		return node.Val
	}
	return -1
}

func traverse(root *TreeNode, cur *int, k int, result *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left, cur, k, result)
	*cur = *cur + 1
	if *cur == k {
		result.Val = root.Val
	}
	traverse(root.Right, cur, k, result)
}
