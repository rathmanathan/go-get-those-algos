// Leetcode Problem:
// https://leetcode.com/problems/concatenation-of-array/description/

// Neetcode 250:
// - Arrays & Hashing
//   - Concatenation of Array

package arrayshashing

func getConcatenation(nums []int) []int {
	n := len(nums)
	result := make([]int, n*2)
	for i, num := range nums {
		result[i], result[i+n] = num, num
	}
	return result
}
