func hasDuplicate(nums []int) bool {
    hashMap := make(map[int]bool)

    for _, val := range nums {
        if hashMap[val] {
            return true
        }
        hashMap[val] = true  
    }

    return false
}
