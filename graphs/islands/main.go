package main

import "fmt"

func dfs(n, m int, grid [][]byte, visited [][]byte) {
	if n < 0 || m < 0 || n >= len(grid) || m >= len(grid[n]) {
		return
	}
	if grid[n][m] == 0 {
		return
	}
	if visited[n][m] == 0 {
		visited[n][m] = 1                                       // mark the current n, m as visited
		if m > 0 && grid[n][m-1] == 1 && visited[n][m-1] == 0 { // check left
			dfs(n, m-1, grid, visited)
		}
		if n > 0 && grid[n-1][m] == 1 && visited[n-1][m] == 0 { // check top
			dfs(n-1, m, grid, visited)
		}
		if m < len(grid[n])-1 && grid[n][m+1] == 1 && visited[n][m+1] == 0 { //check right
			dfs(n, m+1, grid, visited)
		}
		if n < len(grid)-1 && grid[n+1][m] == 1 && visited[n+1][m] == 0 {
			dfs(n+1, m, grid, visited)
		}
	}
	return
}

func numIslands(grid [][]byte) int {
	visited := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]byte, len(grid[i]))
	}
	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if visited[i][j] == 0 && grid[i][j] == 1 {
				islands++
				dfs(i, j, grid, visited)
			}
		}
	}
	return islands
}

func main() {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(numIslands(grid))
}
