package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		changedInput []int
	}{
		{
			name:         "0, 1, 0, 3, 12",
			input:        []int{0, 1, 0, 3, 12},
			changedInput: []int{1, 3, 12, 0, 0},
		},
		{
			name:         "0",
			input:        []int{0},
			changedInput: []int{0},
		},
		{
			name:         "0,1,0,3,12",
			input:        []int{0, 1, 0, 3, 12},
			changedInput: []int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.input)
			assert.Equal(t, tt.changedInput, tt.input)
		})
	}
}
