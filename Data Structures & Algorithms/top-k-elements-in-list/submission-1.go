func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	var sortedValues [][]int
	var result []int

	for _, val := range nums {
		hashMap[val]++
	}

	for key, val := range hashMap {
		sortedValues = append(sortedValues, []int{val, key})
	}

	sort.Slice(sortedValues, func(i, j int) bool {
		return sortedValues[i][0] > sortedValues[j][0]
	})

	count := 0
	for k > count {
		value := sortedValues[count][1]
		result = append(result, value)
		count++
	}

	return result
}
