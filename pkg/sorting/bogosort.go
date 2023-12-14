package sorting

import "github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"

func Bogosort(arr []int) []int {
	for !utils.IsSorted(arr) {
		utils.Shuffle(&arr)
	}
	return arr
}
