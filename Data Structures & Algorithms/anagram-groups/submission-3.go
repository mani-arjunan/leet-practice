func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)

	for _, val := range strs {
		var tuple [26]int
		for _, value := range val {
			tuple[value - 'a']++
		}
		hashMap[tuple] = append(hashMap[tuple], val)
	}

	var result [][]string

	for _, val := range hashMap {
		result = append(result, val)
	}
	return result
}
