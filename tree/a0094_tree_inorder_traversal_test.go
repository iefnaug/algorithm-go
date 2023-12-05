package tree

import (
	"fmt"
	"strconv"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
中序遍历顺序 - 左中右
*/
var ret []int

func inorderTraversal(root *TreeNode) []int {
	inorder(root)
	return ret
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	ret = append(ret, root.Val)
	inorder(root.Right)
}

/*
迭代
*/
func inorderTraversal2(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		root = root.Right
	}
	return ret
}

func Test01(t *testing.T) {
	n7 := TreeNode{7, nil, nil}
	n6 := TreeNode{6, nil, nil}
	n5 := TreeNode{5, nil, nil}
	n4 := TreeNode{4, nil, nil}
	n3 := TreeNode{3, &n6, &n7}
	n2 := TreeNode{2, &n4, &n5}
	n1 := TreeNode{1, &n2, &n3}
	//ret := inorderTraversal(&n1)
	ret := inorderTraversal2(&n1)
	for _, v := range ret {
		fmt.Print(strconv.Itoa(v) + ", ")
	}
}

func Test02(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr[0 : len(arr)-1])
}
