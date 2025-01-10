package arrayshashing

func areMapsEqual(s map[string]int, t map[string]int) bool {
	for k, v := range s {
		if t[k] != v {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var s_map, t_map = make(map[string]int), make(map[string]int)

	for i := range s {
		s_map[string(s[i])]++
		t_map[string(t[i])]++
	}

	return areMapsEqual(s_map, t_map)

}
