package array

import (
	"fmt"
	"math"
	"testing"
)

/*
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。

如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。



示例 1：

输入：nums = [1,2,3,4,5]
输出：true
解释：任何 i < j < k 的三元组都满足题意
示例 2：

输入：nums = [5,4,3,2,1]
输出：false
解释：不存在满足题意的三元组
示例 3：

输入：nums = [2,1,5,0,4,6]
输出：true
解释：三元组 (3, 4, 5) 满足题意，因为 nums[3] == 0 < nums[4] == 4 < nums[5] == 6


提示：

1 <= nums.length <= 5 * 10^5
-2^31 <= nums[i] <= 2^31 - 1


进阶：你能实现时间复杂度为 O(n) ，空间复杂度为 O(1) 的解决方案吗？
*/

/*
动态规划
*/
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	length := len(nums)
	leftMin, rightMax := make([]int, length), make([]int, length)
	leftMin[0] = nums[0]
	rightMax[length-1] = nums[length-1]
	for i := 1; i < len(nums); i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
	}
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}
	for i := 1; i <= length-2; i++ {
		if nums[i] > leftMin[i-1] && nums[i] < rightMax[i+1] {
			return true
		}
	}
	return false
}

func max(i int, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i int, j int) int {
	if i <= j {
		return i
	}
	return j
}

/*
贪心，迭代过程中保证 x<y
*/
func increasingTriplet2(nums []int) bool {
	length := len(nums)
	if length < 3 {
		return false
	}
	x, y := nums[0], math.MaxInt32
	for i := 1; i < length; i++ {
		if nums[i] > y {
			return true
		} else if nums[i] > x {
			y = nums[i]
		} else {
			x = nums[i]
		}
	}
	return false
}

func TestIncreasingTriplet(t *testing.T) {
	nums := []int{
		1, 10, 9, 8, 7, 6,
	}
	exist := increasingTriplet(nums)
	fmt.Println(exist)
}
