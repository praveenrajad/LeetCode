func containsDuplicate(nums []int) bool {
    m := make(map[int]struct{}, len(nums))
    for _, num := range nums {
        if _, ok := m[num]; ok{
            return true
        }
        m[num] = struct{}{}
    }
    return false
}


/*Lessons Learnt

1. make with know size is faster
2. using struct {} is efficient as zero memory is used
3. to add a empty struc{} use struct{}{}

*/

