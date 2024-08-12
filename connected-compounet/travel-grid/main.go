package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Create a local random generator using NewSource
	//zrng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random grid
	grid := generateRandomGrid(5, 5)
	printGrid(grid)

	// Define the number of rows and columns
	m := len(grid)    // Number of rows
	n := len(grid[0]) // Number of columns

	// Initialize visited array
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// Directions: up, down, left, right
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// DFS function
	var dfs func(r, c int)

	dfs = func(r, c int) {
		// Mark current cell as visited
		visited[r][c] = true
		fmt.Printf("Visited cell: (%d, %d)\n", r, c)

		// Explore adjacent cells (up, down, left, right)
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			fmt.Printf("nr: %d - nc %d\n", nr, nc)
			// Check if the adjacent cell is within the grid bounds and is unvisited land
			if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == '1' && !visited[nr][nc] {
				// Recursive call to DFS for the adjacent cell
				dfs(nr, nc)
			}
		}
	}

	// Perform DFS traversal from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				fmt.Printf("\nStarting DFS traversal from cell: (%d, %d)\n", i, j)
				dfs(i, j)
			}
		}
	}
}

// Function to generate a random grid
func generateRandomGrid(m, n int) [][]byte {
	grid := make([][]byte, m)
	for i := range grid {
		grid[i] = make([]byte, n)
		for j := range grid[i] {
			if rand.Intn(2) == 0 {
				grid[i][j] = '0' // Water
			} else {
				grid[i][j] = '1' // Land
			}
		}
	}
	return grid
}

// Function to print the grid
func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}
