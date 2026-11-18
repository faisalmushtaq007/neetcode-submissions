func twoSum(nums []int, target int) []int {
    if len(nums) < 2 {
		return []int{}
	}
    complements := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        
        if index, found := complements[complement]; found {
            return []int{index, i}
        }
        
        complements[num] = i
    }
    
    return nil
}


// return indices 'i' 'j' such that nums[i] + nums[j] == target; i != j
