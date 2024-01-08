package array

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {

	tests := []struct {
		name        string
		input       []int
		wantLength  int
		expectedArr []int
	}{
		{
			name:        "112",
			input:       []int{1, 1, 2},
			wantLength:  2,
			expectedArr: []int{1, 2},
		},
		{
			name:        "[0,0,1,1,1,2,2,3,3,4]",
			input:       []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			wantLength:  5,
			expectedArr: []int{0, 1, 2, 3, 4},
		},
		{
			name:        "[0,0,0,0,0,0]",
			input:       []int{0, 0, 0, 0, 0, 0},
			wantLength:  1,
			expectedArr: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLength := removeDuplicates(tt.input)
			if gotLength != tt.wantLength {
				t.Errorf("removeDuplicates() returned length %v, want %v", gotLength, tt.wantLength)
			}

			// Check if the modified slice matches the expectedArr
			if !reflect.DeepEqual(tt.input[:gotLength], tt.expectedArr) {
				t.Errorf("removeDuplicates() modified array = %v, want %v", tt.input[:gotLength], tt.expectedArr)
			}
		})
	}
}
