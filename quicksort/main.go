package quicksort

import "fmt"

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			aux := arr[i]
			arr[i] = arr[j]
			arr[j] = aux
			i++
		}
	}

	aux := arr[i]
	arr[i] = arr[high]
	arr[high] = aux

	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		arr, p := partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func QuickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println(QuickSortStart(arr))
}
