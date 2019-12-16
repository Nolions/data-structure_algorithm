package quick_sort

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAscendingQuickSort(t *testing.T) {
	values := []int{4, 1, 7, 9, 2, 5, 3}
	ascending(values, 0, len(values)-1)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 7, 9}, values)
}
