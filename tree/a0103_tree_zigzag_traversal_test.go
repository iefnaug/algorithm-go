package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	queue := []*TreeNode{root}
	l2r := true
	level := 0
	for len(queue) > 0 {
		length := len(queue)
		ret = append(ret, make([]int, length))
		newQueue := make([]*TreeNode, 0)
		for i := 0; i < length; i++ {
			if l2r {
				ret[level][i] = queue[i].Val
			} else {
				ret[level][i] = queue[length-1-i].Val
			}
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
		}
		queue = newQueue
		level++
		l2r = !l2r
	}
	return ret
}
