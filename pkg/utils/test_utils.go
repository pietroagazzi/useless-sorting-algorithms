package utils

import (
	"reflect"
	"testing"
)

type SortTest struct {
	Name string
	Arr  []int
	Want []int
}

func RunSortTests(t *testing.T, sortFunc func([]int) []int, tests []SortTest) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := sortFunc(tt.Arr); !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("%s() = %v, want %v", tt.Name, got, tt.Want)
			}
		})
	}
}
