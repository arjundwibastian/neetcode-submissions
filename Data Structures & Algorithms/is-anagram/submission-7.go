func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for _ , v := range s {
		index := strings.Index(t, string(v))
		if index == -1 {
			return false
		}
		t = t[:index] + t[index+1:]
	}
	if t != "" {
		return false
	}
	return true
}
