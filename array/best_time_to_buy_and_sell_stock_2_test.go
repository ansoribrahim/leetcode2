package array

import "testing"

func Test_maxProfit(t *testing.T) {

	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "[7,1,5,3,6,4]",
			input: []int{7, 1, 5, 3, 6, 4},
			want:  7,
		},
		{
			name:  "[1,2,3,4,5,6]",
			input: []int{1, 2, 3, 4, 5, 6},
			want:  5,
		},
		{
			name:  "[7,6,4,3,1]",
			input: []int{7, 6, 4, 3, 1},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.input); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
