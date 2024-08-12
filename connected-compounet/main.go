package main

import "fmt"

func main() {

	adjList := map[int][]int{
		1: {4},
		4: {1, 5},
		5: {4, 7},
		7: {5, 8},
		8: {7},
	}

	v := make(map[int]bool)

	var dfs func(vertice int)
	dfs = func(vertice int) {
		v[vertice] = true
		fmt.Printf("%d ", vertice)

		for _, edge := range adjList[vertice] {
			if !v[edge] {
				dfs(edge)
			}
		}
	}

	// Iterate over all vertices
	for vertice := range adjList {
		if !v[vertice] {
			fmt.Print("Connected component:")
			dfs(vertice)
			fmt.Println()
		}
	}
}
