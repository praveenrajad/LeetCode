func longestConsecutive(nums []int) int {
		m := make(map[int]bool, len(nums))
		for _, num := range nums {
			m[num] = true
		}
		long := 0
		for key := range m {
			if m[key-1] {
				continue
			}
			curr := 1
			key++
			for m[key] {
				curr++
				key++
			}
			if curr > long {
				long = curr
			}
		}
	return long
}