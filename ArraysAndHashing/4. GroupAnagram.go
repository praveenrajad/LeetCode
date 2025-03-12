func groupAnagram( strings []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strings {
		runes := []rune(str)
		sort.Slice(runes, func (i, j int) bool {
			return runes[i] < runes[j]
		})
		m[string(runes)] = append(m[string(runes)], str)
	}
	out := make ([][]string, 0, len(m))
	for _, val := range m {
		out = append(out, val)
	}
	return out
}