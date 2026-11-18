func hasDuplicate(nums []int) bool {
    seen := make(map[int]struct{}) 
    
    for _, value := range nums {
        if _, exists := seen[value]; exists {
            return true
        }
        seen[value] = struct{}{}
    }
    return false
}