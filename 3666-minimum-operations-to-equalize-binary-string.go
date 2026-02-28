package main

func main() {
	minOperations("110", 1)
}

func minOperations(s string, k int) int {
	n := len(s)
	z := 0
	for _, c := range s {
		if c == '0' {
			z++
		}
	}
	o := n - z // count of ones

	if z == 0 {
		return 0
	}

	for i := 1; i <= n; i++ {
		total := i * k

		// Parity check: extra flips beyond z must be even so they cancel out
		if (total-z)%2 != 0 {
			continue
		}
		if total < z {
			continue
		}

		if i%2 == 1 {
			// Odd i: each zero flipped ≤ i times (odd ✓), each one ≤ i-1 times (even ✓)
			if total <= z*i+o*(i-1) {
				return i
			}
		} else {
			// Even i: each zero flipped ≤ i-1 times (odd ✓), each one ≤ i times (even ✓)
			if total <= z*(i-1)+o*i {
				return i
			}
		}
	}

	return -1
}
