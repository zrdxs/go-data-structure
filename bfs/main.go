package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//coins := []int{1, 2, 3, 4, 5}
	//coins := []int{1, 2, 5}
	//nums := []int{1, 2, 3}
	//nums := []int{4, 2, 1}
	nums := []int{2, 7, 9, 3, 1}

	//bfs(coins, 11)

	//combinationSum4(nums, 4)

}

func bfs(coins []int, amount int) int {

	queue := make([]int, len(coins))
	queue = append(queue, amount)
	visited := make(map[int]int)
	steps := 0
	for len(queue) != 0 {
		queue2 := make([]int, len(coins))
		for _, vertex := range queue {
			// pesquisar o pq vertex
			fmt.Printf("O tamanho dessa fila é %d com esses valores %v \n", len(queue), queue)
			for _, coin := range coins {
				// vertex = nó -> confirmar
				nextVertex := vertex - coin

				if nextVertex == 0 {
					return steps + 1
				}
				if nextVertex < 0 {
					break
				}
				if _, ok := visited[nextVertex]; ok {
					continue
				}

				visited[nextVertex] = 1
				fmt.Printf("Adicionando o vertex %d na Fila: \n", nextVertex)
				queue2 = append(queue2, nextVertex)
			}
		}
		steps++
		fmt.Printf("Rodei essa fila pela %d° vez \n", steps)
		queue = queue2
	}

	return -1
}

/* func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	queue := make([]int, 0, len(coins))
	queue = append(queue, amount)

	visited := make(map[int]int)
	steps := 0
	sort.Ints(coins)
	for len(queue) != 0 {
		queue2 := make([]int, 0, len(coins))
		for _, vertex := range queue {
			fmt.Println(queue)
			// add all neighbouring vertices in visited
			for _, edge := range coins {
				nextVertex := vertex - edge
				if nextVertex == 0 {
					return steps + 1
				}
				if nextVertex < 0 {
					break
				}
				if _, ok := visited[nextVertex]; ok {
					continue
				}
				visited[nextVertex] = 1
				queue2 = append(queue2, nextVertex)
			}
		}
		steps += 1
		queue = queue2
	}
	return -1
} */

func combinationSum4(nums []int, target int) int {

	// look-up table
	combinationCount := make(map[int]int)

	// --------------------------------------------------

	// function variable
	var dfs func(int) int

	dfs = func(wanted int) int {

		// base cases:
		if wanted < 0 {
			// stop condition for negative number
			return 0
		} else if wanted == 0 {
			// stop condition for perfect match
			return 1
		}

		if count, exist := combinationCount[wanted]; exist == true {
			// quick response by look-up table
			fmt.Println("Retornando o: ", wanted, " Com a quantidade: ", count)
			return count
		}

		// general cases
		count := 0

		for _, number := range nums {

			nextWanted := wanted - number

			count += dfs(nextWanted)

		}

		combinationCount[wanted] = count
		//fmt.Println("Salvei: ", wanted, " Com a quantidade: ", count)
		return count
	}
	// --------------------------------------------------

	return dfs(target)
}

func bottomUp(nums []int, target int) int {
	// table records solutions for every branch in the decision tree
	// a decision to use a coin denomination creates a branch in the tree
	table := make([]int, target+1)
	table[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			// if the difference == 0, we add 1 (leaf of branch)
			// if the difference > 0, we add the recorded solution for that subproblem (node in branch)
			if i-num >= 0 {
				table[i] += table[i-num]
			}
		}
	}

	json2, _ := json.MarshalIndent(table, "", "  ")
	fmt.Println(string(json2))

	return table[target]
}
