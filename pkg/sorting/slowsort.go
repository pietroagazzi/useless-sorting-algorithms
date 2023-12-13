package sorting

import (
	"math"

	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"
)

func Slowsort(arr []int, start_idx, end_idx int) []int {
	if start_idx >= end_idx {
		return arr
	}

	mid_idx := int(math.Floor(float64(start_idx+end_idx) / 2))

	arr = Slowsort(arr, start_idx, mid_idx)
	arr = Slowsort(arr, mid_idx+1, end_idx)

	if arr[mid_idx] > arr[end_idx] {
		utils.Swap(&arr, mid_idx, end_idx)
	}

	return Slowsort(arr, start_idx, end_idx-1)
}
