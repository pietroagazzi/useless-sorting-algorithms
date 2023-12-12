package sorting

import (
	"math/rand"

	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"
)

func Shuffle(arr *[]int) {
	for i := len(*arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func Bogosort(arr []int) []int {
	for !utils.IsSorted(arr) {
		Shuffle(&arr)
	}
	return arr
}
