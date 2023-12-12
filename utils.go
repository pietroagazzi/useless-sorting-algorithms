package main

import "math/rand"

func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}

	return true
}

func Shuffle(arr *[]int) {
	for i := len(*arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}