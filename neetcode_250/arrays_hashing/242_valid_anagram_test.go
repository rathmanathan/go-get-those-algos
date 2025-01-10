package arrayshashing

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"listen", "silent", true},
		{"hello", "billion", false},
		{"", "", true},
		{"a", "a", true},
		{"ab", "ba", true},
		{"abc", "cba", true},
		{"abcd", "dcba", true},
		{"abcd", "abce", false},
	}

	for _, tc := range tests {
		t.Run(tc.s+"_"+tc.t, func(t *testing.T) {
			got := isAnagram(tc.s, tc.t)
			if got != tc.want {
				t.Fatalf("isAnagram(%v, %v) = %v; want %v", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
