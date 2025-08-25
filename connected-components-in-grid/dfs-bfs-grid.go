package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{1, 1, 0, 0},
		{1, 0, 1, 0},
		{0, 0, 1, 1},
	}

	rows, cols := len(grid), len(grid[0])
	vistiedGraph := make([][]bool, rows)
	for i := range vistiedGraph {
		vistiedGraph[i] = make([]bool, cols)
	}

	// DFS traversal
	for i := range grid {
		for j := range grid[0] {
			if !vistiedGraph[i][j] && grid[i][j] == 1 {
				dfs(grid, vistiedGraph, i, j)
			} else {
				fmt.Printf("already visited [%d],[%d] \n", i, j)
			}
		}
	}

	fmt.Println(vistiedGraph)

	// reset visited
	for i := range vistiedGraph {
		for j := range vistiedGraph[i] {
			vistiedGraph[i][j] = false
		}
	}

	fmt.Println("\nBFS from (0,0):")
	for i := range grid {
		for j := range grid[0] {
			if !vistiedGraph[i][j] && grid[i][j] == 1 {
				bfs(grid, vistiedGraph, Point{i, j})
			} else {
				fmt.Printf("already visited [%d],[%d] \n", i, j)
			}
		}
	}
	fmt.Println(vistiedGraph)
}

var dirs = [][]int{
	{1, 0},  // down
	{-1, 0}, // up
	{0, 1},  // right
	{0, -1}, // left
}

func dfs(grid [][]int, visited [][]bool, r int, c int) {
	rows, cols := len(grid), len(grid[0])

	if r < 0 || c < 0 || r >= rows || c >= cols {
		return
	}

	if visited[r][c] || grid[r][c] == 0 {
		return
	}

	visited[r][c] = true
	fmt.Printf("Visiting: (%d,%d)\n", r, c)

	for _, d := range dirs {
		dfs(grid, visited, r+d[0], c+d[1])
	}
}

type Point struct {
	r int
	c int
}

func bfs(grid [][]int, visited [][]bool, p Point) {
	rows, cols := len(grid), len(grid[0])

	queue := []Point{p}
	visited[p.r][p.c] = true // mark start

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		fmt.Printf("Visiting: (%d,%d)\n", p.r, p.c)

		for _, d := range dirs {
			nr := p.r + d[0]
			nc := p.c + d[1]
			if nr >= 0 && nc >= 0 && nr < rows && nc < cols &&
				!visited[nr][nc] && grid[nr][nc] == 1 {
				visited[nr][nc] = true
				queue = append(queue, Point{nr, nc})
			}
		}
	}
}
