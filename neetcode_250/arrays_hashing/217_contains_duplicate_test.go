package arrayshashing

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1}, true},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			got := containsDuplicate(tc.nums)
			if got != tc.want {
				t.Fatalf("containsDuplicate(%v) = %v; want %v", tc.nums, got, tc.want)
			}
		})
	}
}
