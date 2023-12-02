package linkedlist

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

进阶：你能否设计一个时间复杂度 O(m + n) 、仅用 O(1) 内存的解决方案
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	hashMap := make(map[*ListNode]int)
	for headA != nil {
		hashMap[headA] = headA.Val
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := hashMap[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ha, hb := headA, headB
	for ha != hb {
		if ha == nil {
			ha = headB
		} else {
			ha = ha.Next
		}
		if hb == nil {
			hb = headA
		} else {
			hb = hb.Next
		}
	}
	return ha
}
