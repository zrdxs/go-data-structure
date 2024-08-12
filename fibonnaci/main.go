package main

func fibTab(n int) int {
	tab := make([]int, n+1)
	tab[1] = 1

	for i := 0; i < n; i++ {
		tab[i+1] += tab[i]
		if i+2 <= n {
			tab[i+2] += tab[i]
		}
	}
	return tab[n]
}

func fibMemo(n int, memo map[int]int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	if n <= 2 {
		return 1
	}
	memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
	return memo[n]
}

func main() {
	println(fibMemo(6, make(map[int]int)))
	println(fibMemo(50, make(map[int]int)))
}
