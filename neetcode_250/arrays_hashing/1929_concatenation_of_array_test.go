package arrayshashing

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
		{[]int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
		{[]int{0, 0, 0}, []int{0, 0, 0, 0, 0, 0}},
		{[]int{5, 6, 7}, []int{5, 6, 7, 5, 6, 7}},
		{[]int{}, []int{}},
	}

	for _, tc := range tests {
		t.Run("Testing getConcatenation", func(t *testing.T) {
			result := getConcatenation(tc.input)
			if !reflect.DeepEqual(result, tc.output) {
				t.Fatalf("Expected %v but got %v", tc.output, result)
			}
		})
	}
}
