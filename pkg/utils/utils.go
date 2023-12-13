package utils

import "math/rand"

// IsSorted check if the array is sorted
func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}

	return true
}

// GenerateRandomArray generate a random array with length n, and the value is in [rangeL, rangeR]
func GenerateRandomArray(n, rangeL, rangeR int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rangeL + rand.Intn(rangeR-rangeL+1)
	}

	return arr
}

func Swap(arr *[]int, e1 int, e2 int) {
	(*arr)[e1], (*arr)[e2] = (*arr)[e2], (*arr)[e1]
}
