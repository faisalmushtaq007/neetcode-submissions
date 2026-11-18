func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    charCount := make(map[rune]int)

    for _, r := range s {
        charCount[r]++
    }

    for _, r := range t {
        charCount[r]--
        if charCount[r] < 0 {
            return false 
        }
    }

    return true
}
