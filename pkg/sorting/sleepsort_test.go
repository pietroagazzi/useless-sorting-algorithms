package sorting_test

import (
	"testing"

	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/sorting"
	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"
)

func TestSleepsort(t *testing.T) {
	utils.RunSortTests(t, utils.AdaptSleepsort(sorting.Sleepsort), sorting.Tests)
}
