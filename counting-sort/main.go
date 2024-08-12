package main

import "fmt"

func countingSort(nums []int) []int {

	max := nums[0]
	min := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	counter := make([]int, max-min+1)

	for _, num := range nums {
		counter[num-min]++
	}

	index := 0
	for i, c := range counter {
		for c > 0 {
			nums[index] = i + min
			index++
			c--
		}
	}

	return nums
}

func main() {

	nums := []int{4, 2, -3, 6, -1, 0, 2, -3}
	nums2 := []int{4, 2, 1, 1, 1, 0, 2, 13}

	fmt.Println(countingSort(nums))
	fmt.Println(countingSort(nums2))
}
