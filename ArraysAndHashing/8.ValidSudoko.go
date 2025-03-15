func validSudoku(b [][]byte) bool {
	m := len(b)
	n := len(b[0])
	for i := 0; i < m; i++ {
		s := make(map[byte]bool, m)
		for j := 0; j < n; j++ {
			if b[i][j] == '.' {
				continue
			}
			if s[b[i][j]]  {
				return false
			}
			s[b[i][j]]  = true
		}
	}

	for j := 0; j < n; j++ {
		s := make(map[byte]bool, m)
		for i := 0; i < m; i++ {
			if b[i][j] == '.' {
				continue
			}
			if s[b[i][j]] {
				return false
			}
			s[b[i][j]] = true
		}
	}

	for k:=0; k < m; k++ {
		is := (k/3)*3
		js := (k%3)*3
		s := make(map[byte]bool, m)
		for i := is; i < is+3; i++ {
			for j := js; j < js+3; j++ {
				if b[i][j] == '.' {
					continue
				}
				if s[b[i][j]]  {
					return false
				}
				s[b[i][j]]  = true
			}
		}
	}
	return true
}