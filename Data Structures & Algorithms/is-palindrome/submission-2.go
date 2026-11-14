func isPalindrome(s string) bool {
	chars := []rune(s)

    left := 0
    right := len(chars) - 1

    for left < right {
        l := chars[left]
        r := chars[right]

        // Skip non-alphanumeric characters
        if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
            left++
            continue
        }

        if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
            right--
            continue
        }

        // Compare lowercase characters
        if unicode.ToLower(l) != unicode.ToLower(r) {
            return false
        }

        left++
        right--
    }

    return true
}
