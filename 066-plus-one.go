func plusOne(digits []int) []int {
	addition := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 {
			addition = 1
		}
		val := digits[i] + addition
		if val > 9 {
			digits[i] = val % 10
			addition = 1
		} else {
			digits[i] = val
			addition = 0
		}
	}
	res := make([]int, 0, len(digits)+1)
	if addition != 0 {
		res = append(res, addition)
	}
	res = append(res, digits...)
	return res
}
