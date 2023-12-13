package sorting_test

import (
	"testing"

	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/sorting"
	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"
)

func TestStalinsort(t *testing.T) {
	utils.RunSortTests(t, sorting.Stalinsort, sorting.Tests)
}
