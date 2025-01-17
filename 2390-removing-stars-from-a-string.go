func removeStars(s string) string {
	stack := make([]rune, 0, len(s))
	for _, v := range []rune(s) {
		if v == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}
