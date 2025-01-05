package module01

import (
	"math"
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	value = Reverse(value)
	var char_map = map[string]int{"A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15}
	var dec int
	for ind, char := range value {
		multiplier := int(math.Pow(float64(base), float64(ind)))
		var num int
		if char >= 65 {
			num = char_map[string(char)]
		} else {
			num, _ = strconv.Atoi(string(char))
		}
		dec += num * multiplier
	}
	return dec
}
