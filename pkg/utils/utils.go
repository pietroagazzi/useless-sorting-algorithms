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

type Sorter func([]int) []int

func AdaptSlowsort(f func([]int, int, int) []int) Sorter {
	return func(arr []int) []int {
		return f(arr, 0, len(arr)-1)
	}
}

// AdaptSleepsort adapts a function that works on uint slices to work on int slices
func AdaptSleepsort(f func([]uint) []uint) Sorter {
	return func(arr []int) []int {
		uintArr := make([]uint, len(arr))
		for i, v := range arr {
			uintArr[i] = uint(v)
		}

		result := f(uintArr)

		intArr := make([]int, len(result))
		for i, v := range result {
			intArr[i] = int(v)
		}

		return intArr
	}
}
