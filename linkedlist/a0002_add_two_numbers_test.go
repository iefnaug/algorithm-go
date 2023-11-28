package linkedlist

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。


提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	dummyNode := &ListNode{}
	cur := dummyNode
	for l1 != nil && l2 != nil {
		tmp := l1.Val + l2.Val + add
		if tmp >= 10 {
			add = 1
			tmp -= 10
		} else {
			add = 0
		}
		cur.Next = &ListNode{tmp, nil}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		tmp := l1.Val + add
		if tmp >= 10 {
			add = 1
			tmp -= 10
		} else {
			add = 0
		}
		cur.Next = &ListNode{tmp, nil}
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		tmp := l2.Val + add
		if tmp >= 10 {
			add = 1
			tmp -= 10
		} else {
			add = 0
		}
		cur.Next = &ListNode{tmp, nil}
		cur = cur.Next
		l2 = l2.Next
	}
	if add > 0 {
		cur.Next = &ListNode{add, nil}
		cur = cur.Next
	}
	return dummyNode.Next
}
