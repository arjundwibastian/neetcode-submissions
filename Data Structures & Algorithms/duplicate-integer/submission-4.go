func hasDuplicate(nums []int) bool {
    maps := make(map[int]int)

	for i, v := range nums {
		if _, ok := maps[v]; ok {
			return true
		} else {
			maps[v] = i
		}
	}
	return false
}
