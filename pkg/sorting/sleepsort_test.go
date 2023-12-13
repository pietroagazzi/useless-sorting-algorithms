package sorting_test

import (
	"reflect"
	"testing"

	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/sorting"
)

func TestSleepsort(t *testing.T) {
    tests := []struct {
        name string
        arr  []uint
        want []uint
    }{
        {
            name: "Test 1",
            arr:  []uint{5, 2, 1, 3, 4},
            want: []uint{1, 2, 3, 4, 5},
        },
        {
            name: "Test 2",
            arr:  []uint{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
            want: []uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := sorting.SleepSort(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SleepSort() = %v, want %v", got, tt.want)
			}
        })
    }
}