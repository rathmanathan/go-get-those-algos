package arrayshashing

func containsDuplicate(nums []int) bool {
	var num_hash = make(map[int]bool)
	for _, num := range nums {
		_, exists := num_hash[num]
		if exists {
			return true
		}
		num_hash[num] = true
	}
	return false
}
