func majorityElement(nums []int) int {
	elementLength := make(map[int]int)
	maxKey, maxLen := 0, 0

	for _, value := range nums {
		elementLength[value]++
		if elementLength[value] > maxLen {
			maxLen = elementLength[value]
			maxKey = value
		}
	}

	return maxKey
}
