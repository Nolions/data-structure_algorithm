package bubble_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestBubbleSort ...
func TestBubbleAscendingSort(t *testing.T) {
	values := []int{4, 1, 7, 9, 2, 5, 3}

	ascending(values)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 7, 9}, values)
}

func TestBubbleDescendingSort(t *testing.T) {
	values := []int{4, 1, 7, 9, 2, 5, 3}

	descending(values)
	assert.Equal(t, []int{9, 7, 5, 4, 3, 2, 1}, values)
}
