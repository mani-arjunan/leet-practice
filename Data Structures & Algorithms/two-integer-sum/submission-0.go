func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)

	for i, val := range nums {
		diff := target - val
		value, exists := hashMap[diff]
		if exists {
			return []int{value, i}
		} else {
			hashMap[val] = i
		}
	}
		
	return []int{}
}
