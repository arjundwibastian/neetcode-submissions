func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	b := []byte(t)

	for _,v := range s {
		target := byte(v)
		for in, ch := range b {
			if ch == target {
				b = append(b[:in], b[in+1:]...)
				break
			}

		}
		
	}
	if len(b) != 0 {
		return false
	}
	return true
}
