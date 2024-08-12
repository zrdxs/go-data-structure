package main

import "fmt"

func main() {

	edges := [][]int{{1, 2}, {1, 3}, {2, 4}}

	fmt.Println(IsNonCycle(edges))

}

func IsNonCycle(edges [][]int) bool {
	adjList := make(map[int][]int)

	for _, value := range edges {
		node1, node2 := value[0], value[1]

		_, ok := adjList[node1]
		if !ok {
			adjList[node1] = []int{node2}
		} else {
			adjList[node1] = append(adjList[node1], node2)
		}
	}

	fmt.Println(adjList)

	visited := make(map[int]int)

	for node := range adjList {
		if visited[node] == 1 {
			continue
		}

		if !dfs(node, adjList, &visited) {
			return false
		}
	}

	return true
}

// deep first search
func dfs(node int, adjList map[int][]int, visited *map[int]int) bool {

	neighbourArr, ok := adjList[node]
	if !ok {
		return true
	}

	if (*visited)[node] == -1 {
		return false
	}

	if (*visited)[node] == 1 {
		return true
	}

	(*visited)[node] = -1
	fmt.Println("Nó: ", node, "Setou -1")
	for _, neighbour := range neighbourArr {
		if !dfs(neighbour, adjList, visited) {
			return false
		}
	}

	(*visited)[node] = 1
	fmt.Println("Nó: ", node, "Setou 1")
	return true
}
