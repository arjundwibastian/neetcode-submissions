func firstUniqChar(s string) int {
	mapper := make(map[rune]int)
	for _, v := range s {
		mapper[v]++
	}
	for i, v := range s {
		if mapper[v] == 1 {
			return i
		}
	}
	return -1
}
