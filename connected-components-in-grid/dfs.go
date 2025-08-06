package main

import "fmt"

var deltaRow = []int{-1, 0, 1, 0} // Up, Right, Down, Left
var deltaCol = []int{0, 1, 0, -1}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	n := len(grid)
	m := len(grid[0])
	count := 0

	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if grid[row][col] == '1' {
				count++
				dfs(row, col, grid, n, m)
			}
		}
	}

	return count
}

func dfs(row, col int, grid [][]byte, n, m int) {
	// Mark current cell as visited
	grid[row][col] = '*'

	// Explore 4 directions
	for i := 0; i < 4; i++ {
		nRow := row + deltaRow[i]
		nCol := col + deltaCol[i]

		if nRow >= 0 && nRow < n && nCol >= 0 && nCol < m && grid[nRow][nCol] == '1' {
			dfs(nRow, nCol, grid, n, m)
		}
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println(numIslands(grid)) // Output: 3
}
