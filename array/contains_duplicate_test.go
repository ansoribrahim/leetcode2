package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "[1,2,3,1]",
			input: []int{1, 2, 3, 1},
			want:  true,
		},
		{
			name:  "[1,2,3,4]",
			input: []int{1, 2, 3, 4},
			want:  false,
		},
		{
			name:  "[1,1,1,3,3,4,3,2,4,2]",
			input: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, containsDuplicate(tt.input), "containsDuplicate(%v)", tt.input)
		})
	}
}
