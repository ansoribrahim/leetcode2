package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotateImage(t *testing.T) {
	tests := []struct {
		name           string
		matrix         [][]int
		expectedMatrix [][]int
	}{
		{
			name: "normal",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expectedMatrix: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "normal",
			matrix: [][]int{
				{2, 29, 20, 26, 16, 28},
				{12, 27, 9, 25, 13, 21},
				{32, 33, 32, 2, 28, 14},
				{13, 14, 32, 27, 22, 26},
				{33, 1, 20, 7, 21, 7},
				{4, 24, 1, 6, 32, 34},
			},
			expectedMatrix: [][]int{
				{4, 33, 13, 32, 12, 2},
				{24, 1, 14, 33, 27, 29},
				{1, 20, 32, 32, 9, 20},
				{6, 7, 27, 2, 25, 26},
				{32, 21, 22, 28, 13, 16},
				{34, 7, 26, 14, 21, 28},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateImage(tt.matrix)
			assert.Equal(t, tt.expectedMatrix, tt.matrix)
		})
	}
}
