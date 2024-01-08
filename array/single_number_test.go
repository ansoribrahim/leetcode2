package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "[2,2,1]",
			input: []int{2, 2, 1},
			want:  1,
		},
		{
			name:  "[4,1,2,1,2]",
			input: []int{4, 1, 2, 1, 2},
			want:  4,
		},
		{
			name:  "[1]",
			input: []int{1},
			want:  1,
		},
		{
			name:  "[476,1,298,1,298]",
			input: []int{476, 1, 298, 1, 298},
			want:  476,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, singleNumberUsingXor(tt.input), "singleNumber(%v)", tt.input)
			assert.Equalf(t, tt.want, singleNumber(tt.input), "singleNumber(%v)", tt.input)
		})
	}
}
