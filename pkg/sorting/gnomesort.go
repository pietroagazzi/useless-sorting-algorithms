package sorting

import "github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"

func Gnomesort(arr []int) []int {
	var pos int

	for pos < len(arr) {
		if pos == 0 || arr[pos] >= arr[pos-1] {
			pos++
		} else {
			utils.Swap(&arr, pos, pos-1)
			pos--
		}
	}

	return arr
}
