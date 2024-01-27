package array

import (
	"fmt"
	"testing"
)

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。
*/

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				//将所有相邻的1替换为0
				change2Zero(grid, i, j)
				count++
			}
		}
	}
	return count
}

func change2Zero(grid [][]byte, i int, j int) {
	x := len(grid)
	y := len(grid[0])
	if i < 0 || i >= x || j < 0 || j >= y {
		return
	}
	if grid[i][j] == 1 {
		grid[i][j] = 0
		change2Zero(grid, i-1, j)
		change2Zero(grid, i+1, j)
		change2Zero(grid, i, j-1)
		change2Zero(grid, i, j+1)
	}
}

func TestNumIslands(t *testing.T) {
	//grid := [][]byte{
	//	{1, 1, 0, 1, 0},
	//	{0, 0, 1, 0, 0},
	//	{1, 1, 0, 1, 0},
	//	{0, 0, 0, 0, 0}}
	//grid := [][]byte{
	//	{1, 1, 1, 1, 0},
	//	{1, 1, 0, 1, 0},
	//	{1, 1, 0, 0, 0},
	//	{0, 0, 0, 0, 0}}
	grid := [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1}}
	islands := numIslands(grid)
	fmt.Println(islands)
}
