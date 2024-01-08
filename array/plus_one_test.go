package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "1, 2, 3",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 4},
		},
		{
			name:  "4, 3, 2, 1",
			input: []int{4, 3, 2, 1},
			want:  []int{4, 3, 2, 2},
		},
		{
			name:  "9",
			input: []int{9},
			want:  []int{1, 0},
		},
		{
			name:  "1, 2, 9",
			input: []int{1, 2, 9},
			want:  []int{1, 3, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, plusOne(tt.input), "plusOne(%v)", tt.input)
		})
	}
}
