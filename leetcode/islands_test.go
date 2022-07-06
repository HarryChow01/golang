package leetcode

import "testing"

func TestIslandNum(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	count := numIslands(grid)
	t.Log("count = ", count)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	// 动态创建二维数组
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && visit[i][j] == false {
				markConnectIsland(grid, visit, i, j)
				count++
			}
		}
	}
	return count
}

func markConnectIsland(grid [][]byte, visitArray [][]bool, row int, column int) {
	if row < 0 || row >= len(grid) {
		return
	}
	if column < 0 || column >= len(grid[0]) {
		return
	}
	if grid[row][column] == '0' || visitArray[row][column] {
		return
	}
	visitArray[row][column] = true

	markConnectIsland(grid, visitArray, row-1, column)
	markConnectIsland(grid, visitArray, row+1, column)
	markConnectIsland(grid, visitArray, row, column-1)
	markConnectIsland(grid, visitArray, row, column+1)
}
