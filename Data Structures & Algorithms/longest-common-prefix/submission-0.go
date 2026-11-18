func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Use the first string as the baseline
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}


// input -> array of string e.g., strs = ["bat","bag","bank","band"]
// return the longest commonn prefix of all strings
// e.g., for the input -> output shpuld be -> "ba"
