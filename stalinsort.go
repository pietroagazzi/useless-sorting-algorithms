package main

func StalinSort(arr []int) []int {
	var sorted []int

	if len(arr) == 0 {
		return sorted
	}

	sorted = arr[:1]

	for i := 1; i < len(arr); i++ {
		if arr[i] >= sorted[len(sorted) - 1] {
			sorted = append(sorted, arr[i])
		}
	}

	return sorted
}