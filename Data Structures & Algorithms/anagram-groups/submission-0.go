func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	total := 0
	hashMap := make(map[string]int)

	for _, val := range s {
		hashMap[string(val)]++
		total++
	}

	for _, val := range t {
		if value, exists := hashMap[string(val)]; exists && value > 0{
			hashMap[string(val)]--
			total--
		} 
	}

	return total == 0
}

func groupAnagrams(strs []string) [][]string {
	length := len(strs)
	seen := make([]bool, length)
	var result [][]string
	hashMap := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		for j := i; j < len(strs); j++ {
			if seen[j] {
				continue
			}
			isAna := isAnagram(strs[i], strs[j])

			if isAna {
				_, exists := hashMap[strs[i]]
				if exists {
					hashMap[strs[i]] = append(hashMap[strs[i]], strs[j])
				} else {
					hashMap[strs[i]] = []string{strs[j]}
				}
				seen[j] = true
			}
		}
	}

	for _, value := range hashMap {
		result = append(result, value)
	}
	return result
}
