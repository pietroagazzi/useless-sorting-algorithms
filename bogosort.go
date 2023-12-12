package main

func Bogosort(arr []int) []int {
	for !IsSorted(arr) {
		Shuffle(&arr)
	}
	return arr
}
