func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
    for i, num := range nums {
        if _, ok := m[target - num]; ok {
            return []int{m[target-num], i}
        }
        m[num]=i
    }
    return []int{}
}