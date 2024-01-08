package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersect(t *testing.T) {

	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "normal",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2, 2},
		},
		{
			name:  "normal",
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{9, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, intersect(tt.nums1, tt.nums2), "intersect(%v, %v)", tt.nums1, tt.nums2)
		})
	}
}
