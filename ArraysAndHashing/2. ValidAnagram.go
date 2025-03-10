func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m := make(map[rune]int, len(s))
    for _, val := range s {
        m[val]++
    }
    for _, val := range t {
        m[val]--
        if m[val] < 0 {
            return false
        }    
    }

    return true
}