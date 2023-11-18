package array

import (
	"testing"
)

/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
*/

/*
时间复杂度: O(n^2)
空间复杂度: O(m+n)
*/
func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func TestSetZeros(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{1, 0, 3},
		{1, 2, 3},
	}
	setZeroes(matrix)
}
