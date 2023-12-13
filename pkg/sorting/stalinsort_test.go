package sorting_test

import (
	"reflect"
	"testing"

	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/sorting"
)

func TestStalinsort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Test 1",
			arr:  []int{5, 2, 1, 3, 4},
			want: []int{5},
		},
		{
			name: "Test 2",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want: []int{9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sorting.Stalinsort(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stalinsort() = %v, want %v", got, tt.want)
			}
		})
	}
}
