package main

import "fmt"

func main() {

	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 0},
	}
	k := 1
	//Output: 6

	//grid := [][]int{{0, 1, 1}, {1, 1, 1}, {1, 0, 0}}
	//k := 1
	// Output: -1

	fmt.Println(gridSmallPath(grid, k))
}

type State struct {
	x, y, obstacles, steps int
}

func gridSmallPath(grid [][]int, k int) int {

	r := len(grid)
	c := len(grid[0])

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited := make([][][]bool, r)
	for i := range visited {
		visited[i] = make([][]bool, c)
		for j := range visited[i] {
			visited[i][j] = make([]bool, k+1)
		}
	}
	visited[0][0][0] = true

	queue := []State{{0, 0, 0, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.x == r-1 && curr.y == c-1 {
			return curr.steps
		}

		for _, dir := range directions {
			nx, ny := curr.x+dir[0], curr.y+dir[1]
			if nx >= 0 && nx < r && ny >= 0 && ny < c {
				nk := curr.obstacles + grid[nx][ny]
				if nk <= k && !visited[nx][ny][nk] {
					visited[nx][ny][nk] = true
					queue = append(queue, State{nx, ny, nk, curr.steps + 1})
				}
			}
		}
	}

	return -1
}

/*
{0, 0, 0}
{1, 1, 0}
{0, 0, 0}
{0, 1, 1}
{0, 0, 0}

*/
