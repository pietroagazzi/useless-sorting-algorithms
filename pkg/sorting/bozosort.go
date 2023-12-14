package sorting

import (
	"math/rand"

	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"
)

func Bozosort(arr []int) []int {
	for !utils.IsSorted(arr) {
		i1, i2 := rand.Intn(len(arr)), rand.Intn(len(arr))

		utils.Swap(&arr, i1, i2)
	}

	return arr
}
