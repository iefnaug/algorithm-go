package tree

/*
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		newQueue := make([]*Node, 0)
		for i := 0; i < len(queue); i++ {
			leftNode := queue[i].Left
			rightNode := queue[i].Right
			if i < len(queue)-1 {
				queue[i].Next = queue[i+1]
			}
			if leftNode != nil {
				newQueue = append(newQueue, leftNode)
			}
			if rightNode != nil {
				newQueue = append(newQueue, rightNode)
			}
		}
		queue = newQueue
	}
	return root
}
