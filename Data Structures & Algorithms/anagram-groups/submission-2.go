func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string)

	for _, val := range strs {
		s := sortString(val)
		hashMap[s] = append(hashMap[s], val)
	}

	var result [][]string
	for _, val := range hashMap {
		result = append(result, val)
	}

	return result
}
