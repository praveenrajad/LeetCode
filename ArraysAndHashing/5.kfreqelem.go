func kTopFreqElements(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))
	for _, num :=range nums {
		freq[num]++
	}
	bucket := make([][]int, len(nums)+1)
	for key, val := range freq {
		bucket[val] = append(bucket[val], key)
	}
	out := make([]int, 0, k)
	for i := len(nums); i > 0; i-- {
		for  _, val := range bucket[i] {
			out = append(out, val )
			k--
			if k == 0 {
				return out
			}
		}
	}
	return out
}