func minPartitions(n string) int {
	max := 0
	for _, c := range n {
		digit := int(c - '0')
		if digit > max {
			max = digit
		}
		if max == 9 {
			break // can't get higher, short-circuit
		}
	}
	return max
}