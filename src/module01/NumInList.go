package module01

func NumInList(list []int, num int) bool {
	for _, val := range list {
		if val == num {
			return true
		}
	}
	return false
}
