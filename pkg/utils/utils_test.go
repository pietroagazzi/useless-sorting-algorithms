package utils_test

import (
	"reflect"
	"testing"

	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"
)

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "Test 1",
			arr:  []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "Test 2",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want: false,
		},
		{
			name: "Test 3",
			arr:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: false,
		},
		{
			name: "Test 4",
			arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: true,
		},
		{
			name: "Test 5",
			arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 0},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsSorted(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
