package linkedlist

import (
	"fmt"
	"testing"
)

/*
给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。

第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。

请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。

你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。

提示:

n ==  链表中的节点数
0 <= n <= 10^4
-10^6 <= Node.val <= 10^6
*/

func oddEvenList(head *ListNode) *ListNode {
	oddStart, evenStart := &ListNode{}, &ListNode{}
	odd, even := oddStart, evenStart
	useOdd := true
	for head != nil {
		if useOdd {
			odd.Next = head
			odd = odd.Next
			useOdd = false
		} else {
			even.Next = head
			even = even.Next
			useOdd = true
		}
		head = head.Next
	}
	odd.Next = evenStart.Next
	return oddStart.Next
}

func oddEvenList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func TestOddEvenList(t *testing.T) {
	l6 := &ListNode{6, nil}
	l5 := &ListNode{5, l6}
	l4 := &ListNode{4, l5}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{2, l3}
	l1 := &ListNode{1, l2}
	list := oddEvenList(l1)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
