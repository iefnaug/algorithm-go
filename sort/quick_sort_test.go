package sort

import (
	"fmt"
	"testing"
)

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

func partition(nums []int, left int, right int) int {
	i, j := left, right
	v := nums[i]
	for {
		for nums[i] <= v && i < right {
			i++
		}
		for nums[j] >= v && j > left {
			j--
		}
		if i >= j {
			break
		}
		swap(nums, i, j)
	}
	swap(nums, left, j)
	return j
}

func swap(nums []int, i int, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func TestQuickSort(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0, 5, 5, 5}
	quickSort(nums, 0, len(nums)-1)
	for _, v := range nums {
		fmt.Println(v)
	}
}
