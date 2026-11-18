func removeElement(nums []int, val int) int {
	k := 0

	for _, value := range nums {
		if value != val {
			nums[k] = value
			k++
		}
	}

	return k
}
