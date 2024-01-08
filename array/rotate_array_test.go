package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		k            int
		expectedNums []int
	}{
		{
			name:         "normal",
			nums:         []int{1, 2, 3, 4, 5, 6, 7},
			k:            3,
			expectedNums: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:         "rotate more than length",
			nums:         []int{1, 2, 3, 4, 5, 6, 7},
			k:            8,
			expectedNums: []int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			name:         "with negative value",
			nums:         []int{-1, -100, 3, 99},
			k:            2,
			expectedNums: []int{3, 99, -1, -100},
		},
		{
			name:         "one value",
			nums:         []int{-1},
			k:            7,
			expectedNums: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.nums, tt.k)
			assert.Equal(t, tt.expectedNums, tt.nums)
		})
	}
}
