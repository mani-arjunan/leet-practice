func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	total := 0
	hashMap := make(map[string]int)

	for _, val := range s {
		v := string(val)
		_, exists := hashMap[v]
		if exists {
			hashMap[v]++
		} else {
			hashMap[v] = 1
		}
		total++
	}
	for _, val := range t {
		v := string(val)
		value, exists := hashMap[v]

		if exists && value > 0{
			hashMap[v]--
			total--
		}
	}

	return total == 0
}
