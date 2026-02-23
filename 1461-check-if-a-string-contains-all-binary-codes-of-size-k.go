func hasAllCodes(s string, k int) bool {
	table := make(map[int64]struct{})
	if k >= len(s) {
		return false
	}

	for l, r := 0, k; r <= len(s); {
		substring := s[l:r]
		l++
		r++

		key, _ := strconv.ParseInt(substring, 2, 64)
		table[key] = struct{}{}
	}

	for i := int64(0); i < (2 << (k - 1)); i++ {
		if _, ok := table[i]; !ok {
			return false
		}
	}
	return true
}