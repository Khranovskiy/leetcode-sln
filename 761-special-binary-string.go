package main

import (
	"fmt"
	"sort"
	"strings"
)

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}

	var parts []string
	balance := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			balance++
		} else {
			balance--
		}

		// Found a top-level special substring s[start : i+1]
		if balance == 0 {
			// Strip outer '1' and '0', recursively optimize inside
			inner := makeLargestSpecial(s[start+1 : i])
			parts = append(parts, "1"+inner+"0")
			start = i + 1
		}
	}

	// Put larger blocks first
	sort.Slice(parts, func(i, j int) bool {
		return parts[i] > parts[j]
	})

	return strings.Join(parts, "")
}

func main() {
	fmt.Println(makeLargestSpecial("11011000"))
}
