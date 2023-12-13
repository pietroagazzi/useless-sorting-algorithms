package sorting_test

import (
	"reflect"
	"testing"

	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/sorting"
	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"
)

// Algorithms to test
var algorithms = []struct {
	name string
	f    func([]int) []int
}{
	{"Bogosort", sorting.Bogosort},
	{"Gnomesort", sorting.Gnomesort},
	{"Sleepsort", utils.AdaptSleepsort(sorting.Sleepsort)},
	{"Slowsort", utils.AdaptSlowsort(sorting.Slowsort)},
	{"Stalinsort", sorting.Stalinsort},
}

type SortTest struct {
	Name string
	Arr  []int
	Want []int
}

var Tests = []SortTest{
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

// TestSorting tests all the sorting algorithms
func TestSorting(t *testing.T) {
	for _, alg := range algorithms {
		for _, test := range Tests {
			defer t.Run(alg.name+" "+test.Name, func(t *testing.T) {
				got := alg.f(test.Arr)
				if !reflect.DeepEqual(got, test.Want) {
					t.Errorf("got %v, want %v", got, test.Want)
				}
			})
		}
	}
}