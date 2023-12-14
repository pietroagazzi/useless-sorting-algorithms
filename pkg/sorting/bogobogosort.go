package sorting

import "github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"

func Bogobogosort(arr []int) []int {
	var index int = 2

	for !utils.IsSorted(arr) {
		Bogosort(arr[:index])
		index++

		if !utils.IsSorted(arr[:index]) {
			utils.Shuffle(&arr)
			index = 2
		}
	}

	return arr
}
