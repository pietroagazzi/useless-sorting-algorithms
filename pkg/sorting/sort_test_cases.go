package sorting

import "github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"

var Tests = []utils.SortTest{
	{
		Name: "Test 1",
		Arr:  []int{5, 2, 1, 3, 4},
		Want: []int{1, 2, 3, 4, 5},
	},
	{
		Name: "Test 2",
		Arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		Want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		Name: "Test 3",
		Arr:  []int{1, 2, 3, 4, 5},
		Want: []int{1, 2, 3, 4, 5},
	},
}
