func getConcatenation(nums []int) []int {
    ans := append([]int{}, nums...)

	ans = append(ans, nums...)
	return ans
}